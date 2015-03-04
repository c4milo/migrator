// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package migrator

import (
	"bytes"
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
	"testing"

	"github.com/c4milo/migrator/migrations"
	"github.com/hooklift/assert"
)

func init() {
	if d := os.Getenv("DEBUG"); d != "" {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
	} else {
		log.SetOutput(ioutil.Discard)
	}
}

func runAndLog(cmd *exec.Cmd) (string, string, error) {
	var stdout, stderr bytes.Buffer

	log.Printf("Executing: %s %v", cmd.Path, cmd.Args[1:])
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()

	stdoutString := strings.TrimSpace(stdout.String())
	stderrString := strings.TrimSpace(stderr.String())

	if _, ok := err.(*exec.ExitError); ok {
		message := stderrString
		if message == "" {
			message = stdoutString
		}

		err = fmt.Errorf("error: %s", message)
	}

	log.Printf("stdout: %s", stdoutString)
	log.Printf("stderr: %s", stderrString)

	// Replace these for Windows, we only want to deal with Unix
	// style line endings.
	returnStdout := strings.Replace(stdout.String(), "\r\n", "\n", -1)
	returnStderr := strings.Replace(stderr.String(), "\r\n", "\n", -1)

	return returnStdout, returnStderr, err
}

func createUser() {
	cmd := exec.Command("createuser", "migrator", "--echo", "--login", "--superuser")
	runAndLog(cmd)
}

func destroyUser() {
	cmd := exec.Command("dropuser", "migrator", "--echo")
	runAndLog(cmd)
}

func createDB() {
	cmd := exec.Command("createdb", "migrator_ci", "--echo", "--owner", "migrator")
	runAndLog(cmd)
}

func destroyDB() {
	db, _ = sql.Open("postgres", "user=migrator dbname=migrator_ci sslmode=disable")
	db.Exec("update pg_database set datallowconn = 'false' where datname = 'migrator_ci'")
	db.Exec("select pg_terminate_backend(pid) from pg_stat_activity where datname = 'migrator_ci'")

	cmd := exec.Command("dropdb", "migrator_ci", "--echo")
	runAndLog(cmd)
}

var db *sql.DB

func TestMain(m *testing.M) {
	actions := []func(){
		createUser,
		createDB,
	}

	for _, actionFunc := range actions {
		actionFunc()
	}

	db, _ = sql.Open("postgres", "user=migrator dbname=migrator_ci sslmode=disable")
	// Run tests
	exitCode := m.Run()
	db.Close()

	// tear down
	dactions := []func(){
		destroyDB,
		destroyUser,
	}

	for _, actionFunc := range dactions {
		actionFunc()
	}

	os.Exit(exitCode)
}

func TestMigrateFromAsset(t *testing.T) {
	m, err := NewMigrator(db, Postgres)
	assert.Ok(t, err)

	err = m.Init()
	assert.Ok(t, err)

	err = m.MigrateFromAsset(migrations.Asset, migrations.AssetDir)
	assert.Ok(t, err)

	row := db.QueryRow("select count(*) from schema_migrations")
	var tm int
	row.Scan(&tm)
	assert.Equals(t, 7, tm)
}

func TestRollback(t *testing.T) {
	m, err := NewMigrator(db, Postgres)
	assert.Ok(t, err)

	wor := db.QueryRow("select to_regclass('tokens')")
	var tt string
	wor.Scan(&tt)
	assert.Equals(t, "tokens", tt)

	err = m.Rollback()
	assert.Ok(t, err)

	row := db.QueryRow("select count(*) from schema_migrations where status=$1", "down")

	var tm int
	row.Scan(&tm)
	assert.Equals(t, 1, tm)

	wor2 := db.QueryRow("select to_regclass('tokens')")
	var tt2 string
	wor2.Scan(&tt2)
	assert.Equals(t, "", tt2)

	TestMigrateFromAsset(t)

	row2 := db.QueryRow("select count(*) from schema_migrations where status=$1", "down")
	var tm2 int
	row2.Scan(&tm2)
	assert.Equals(t, 0, tm2)

	row3 := db.QueryRow("select count(*) from schema_migrations")
	var tm3 int
	row3.Scan(&tm3)
	assert.Equals(t, 7, tm3)

	err = m.RollbackN(3)
	row4 := db.QueryRow("select count(*) from schema_migrations where status=$1", "down")

	var tm4 int
	row4.Scan(&tm4)
	assert.Equals(t, 3, tm4)
}

func TestMigrations(t *testing.T) {
	m, err := NewMigrator(db, Postgres)
	assert.Ok(t, err)

	ms, err := m.Migrations()
	assert.Ok(t, err)
	assert.Equals(t, 7, len(ms))
}
