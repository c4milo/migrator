package migrations

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"reflect"
	"strings"
	"unsafe"
	"os"
	"time"
	"io/ioutil"
	"path"
	"path/filepath"
)

func bindata_read(data, name string) ([]byte, error) {
	var empty [0]byte
	sx := (*reflect.StringHeader)(unsafe.Pointer(&data))
	b := empty[:]
	bx := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	bx.Data = sx.Data
	bx.Len = len(data)
	bx.Cap = bx.Len

	gz, err := gzip.NewReader(bytes.NewBuffer(b))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindata_file_info struct {
	name string
	size int64
	mode os.FileMode
	modTime time.Time
}

func (fi bindata_file_info) Name() string {
	return fi.name
}
func (fi bindata_file_info) Size() int64 {
	return fi.size
}
func (fi bindata_file_info) Mode() os.FileMode {
	return fi.mode
}
func (fi bindata_file_info) ModTime() time.Time {
	return fi.modTime
}
func (fi bindata_file_info) IsDir() bool {
	return false
}
func (fi bindata_file_info) Sys() interface{} {
	return nil
}

var _migrations_migrations_go = "\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00"

func migrations_migrations_go_bytes() ([]byte, error) {
	return bindata_read(
		_migrations_migrations_go,
		"migrations/migrations.go",
	)
}

func migrations_migrations_go() (*asset, error) {
	bytes, err := migrations_migrations_go_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "migrations/migrations.go", size: 0, mode: os.FileMode(420), modTime: time.Unix(1425517563, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _migrations_postgres_0001_create_index_function_if_not_exists_down_sql = "\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x54\xca\x41\x0a\x80\x30\x0c\x44\xd1\xab\x64\xa9\xe0\x0d\x3c\x4c\x28\x6d\x0a\x03\x9a\x6a\x3b\x42\x8e\x2f\x14\x37\x2e\x1f\xff\x97\xde\x2e\xa9\x8f\x67\xa2\xb9\xa0\x8a\x05\x06\x87\xe4\x6e\x89\xa6\xf0\x62\xa1\xa8\xea\x8d\xfa\xa5\x85\xea\xe9\x34\xa1\x05\x37\xc1\x0f\x73\x1f\xf7\x31\xbd\xee\x6f\x00\x00\x00\xff\xff\x77\x74\xd0\xd5\x5e\x00\x00\x00"

func migrations_postgres_0001_create_index_function_if_not_exists_down_sql_bytes() ([]byte, error) {
	return bindata_read(
		_migrations_postgres_0001_create_index_function_if_not_exists_down_sql,
		"migrations/postgres/0001_create-index-function-if-not-exists_down.sql",
	)
}

func migrations_postgres_0001_create_index_function_if_not_exists_down_sql() (*asset, error) {
	bytes, err := migrations_postgres_0001_create_index_function_if_not_exists_down_sql_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "migrations/postgres/0001_create-index-function-if-not-exists_down.sql", size: 94, mode: os.FileMode(420), modTime: time.Unix(1425515484, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _migrations_postgres_0001_create_index_function_if_not_exists_up_sql = "\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x5c\x51\xc1\x8e\xab\x30\x0c\xbc\xe7\x2b\xe6\x50\x89\x57\xe9\x15\x69\xcf\xa8\xdf\x82\xd2\x60\x20\xab\xd4\x61\x93\xd0\x76\xa5\x7e\xfc\x3a\x81\x6d\x59\x4e\x4e\xec\xb1\x3d\x9e\x39\x9d\x90\x46\x1b\xd1\xcf\x6c\x92\xf5\x0c\xed\x9c\xbf\x47\xcc\x11\xc9\xc3\x04\xd2\x89\x60\xb9\xa3\x07\x45\xd8\x5e\xc0\xf4\x8d\xce\x73\x95\x40\x0f\x1b\x93\x5a\x21\x3e\x20\xd0\xe4\xb4\xa1\xf7\xa8\xa5\xd4\x96\xee\xd6\xf6\x2d\xfb\xd4\x96\xa6\x88\x7f\xa9\x65\x7d\x25\x24\x7a\xa4\xff\xb0\x7f\x3e\x05\x1e\xbf\x5c\xf9\x1f\x65\x6c\x9a\x03\x47\xdc\xbc\xed\xa0\x23\x0e\x07\xd5\x91\x71\x3a\x90\x82\xec\x72\x6e\x5d\x50\x46\xdc\x74\x30\xa3\x0e\x8d\x94\xa2\x19\xe9\xaa\x77\xe9\x0b\x0d\x96\x95\xda\xb7\x9d\xb1\xf2\x79\x3e\x51\xb5\x55\x0e\x0b\xa7\x46\x6d\xc7\x9c\x51\x4d\xf3\xc5\x59\x53\x35\x4a\x89\x18\x72\x10\x7e\x0f\x92\x8d\xb2\x93\x1c\x99\x84\x8f\xf2\xe9\x83\xbf\x4a\x98\x86\x56\xd8\xc6\x08\x53\xb2\x9f\xde\xf2\x92\xcd\x23\xe3\x94\x15\x63\x88\x5a\x5c\xe7\x03\xcf\x30\x75\x20\xf7\xaa\x95\x9e\xfb\x48\x81\xf0\xae\x08\x6a\x77\x41\x81\x69\xee\x72\xe0\x9a\xe3\xb4\xc2\x36\xec\x0b\xe4\x98\x0d\x14\x01\xf2\x5b\x2c\x35\xb3\x58\x57\x6d\x5d\x46\x39\x7e\xaf\x4f\x96\x25\x73\x2c\xc5\xad\x22\xb9\x50\x97\xec\x46\xc0\x05\xf6\xf2\xb1\x51\x24\xc4\x6c\x5f\xa2\x12\xfb\x9c\xe6\x61\xd6\x03\x61\x72\xd3\x90\x8d\xbe\x79\xa7\x93\x75\xd4\xfc\x04\x00\x00\xff\xff\xf5\xe7\x82\xd2\x8e\x02\x00\x00"

func migrations_postgres_0001_create_index_function_if_not_exists_up_sql_bytes() ([]byte, error) {
	return bindata_read(
		_migrations_postgres_0001_create_index_function_if_not_exists_up_sql,
		"migrations/postgres/0001_create-index-function-if-not-exists_up.sql",
	)
}

func migrations_postgres_0001_create_index_function_if_not_exists_up_sql() (*asset, error) {
	bytes, err := migrations_postgres_0001_create_index_function_if_not_exists_up_sql_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "migrations/postgres/0001_create-index-function-if-not-exists_up.sql", size: 654, mode: os.FileMode(420), modTime: time.Unix(1425515484, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _migrations_postgres_0002_create_extension_citext_down_sql = "\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xd2\xd5\x55\x28\xcd\xcb\xc9\x4f\x4c\x29\x56\x48\xce\x2c\x49\xad\x28\x51\x00\xe2\xd4\xbc\xe2\xcc\xfc\x3c\xae\x94\xa2\xfc\x02\x04\x57\x21\x33\x0d\xc8\xc9\x2c\x2e\x81\xa9\xb4\x06\x04\x00\x00\xff\xff\xf0\xc4\x1b\xe6\x3c\x00\x00\x00"

func migrations_postgres_0002_create_extension_citext_down_sql_bytes() ([]byte, error) {
	return bindata_read(
		_migrations_postgres_0002_create_extension_citext_down_sql,
		"migrations/postgres/0002_create-extension-citext_down.sql",
	)
}

func migrations_postgres_0002_create_extension_citext_down_sql() (*asset, error) {
	bytes, err := migrations_postgres_0002_create_extension_citext_down_sql_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "migrations/postgres/0002_create-extension-citext_down.sql", size: 60, mode: os.FileMode(420), modTime: time.Unix(1425515484, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _migrations_postgres_0002_create_extension_citext_up_sql = "\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xd2\xd5\x55\xc8\xc9\x4f\x4c\x29\x56\x48\xce\x2c\x49\xad\x28\x51\x00\xe2\xd4\xbc\xe2\xcc\xfc\x3c\xae\xe4\xa2\xd4\xc4\x92\x54\x84\x80\x42\x66\x9a\x42\x5e\x3e\x48\x45\x66\x71\x09\x4c\xbd\x35\x20\x00\x00\xff\xff\x0e\xa3\xe6\x8d\x40\x00\x00\x00"

func migrations_postgres_0002_create_extension_citext_up_sql_bytes() ([]byte, error) {
	return bindata_read(
		_migrations_postgres_0002_create_extension_citext_up_sql,
		"migrations/postgres/0002_create-extension-citext_up.sql",
	)
}

func migrations_postgres_0002_create_extension_citext_up_sql() (*asset, error) {
	bytes, err := migrations_postgres_0002_create_extension_citext_up_sql_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "migrations/postgres/0002_create-extension-citext_up.sql", size: 64, mode: os.FileMode(420), modTime: time.Unix(1425515484, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _migrations_postgres_0003_create_accounts_table_down_sql = "\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x4a\x29\xca\x2f\x50\x28\x49\x4c\xca\x49\x55\xc8\x4c\x53\x48\xad\xc8\x2c\x2e\x29\x56\x48\x4c\x4e\xce\x2f\xcd\x2b\x29\xb6\xe6\x82\x48\x57\x16\x60\x91\x8d\x2f\x2e\x49\x2c\x29\x2d\x8e\x07\xc9\xe2\x57\x08\x51\x01\x08\x00\x00\xff\xff\x14\x37\xe9\x6e\x6a\x00\x00\x00"

func migrations_postgres_0003_create_accounts_table_down_sql_bytes() ([]byte, error) {
	return bindata_read(
		_migrations_postgres_0003_create_accounts_table_down_sql,
		"migrations/postgres/0003_create-accounts-table_down.sql",
	)
}

func migrations_postgres_0003_create_accounts_table_down_sql() (*asset, error) {
	bytes, err := migrations_postgres_0003_create_accounts_table_down_sql_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "migrations/postgres/0003_create-accounts-table_down.sql", size: 106, mode: os.FileMode(420), modTime: time.Unix(1425515484, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _migrations_postgres_0003_create_accounts_table_up_sql = "\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xa4\x54\x4b\x92\xd4\x30\x0c\xdd\xe7\x14\xaa\xde\xf4\x4c\x55\x86\x03\xc0\x96\x0d\xa7\x48\x79\x1c\xa5\x5b\x85\x63\x07\x7f\xfa\xc3\xe9\x91\xad\x38\xd5\x24\xa1\x19\x0a\xad\x6c\x4b\x7a\x4f\x8a\x9e\xf2\xf6\x06\xda\xa3\x8a\x18\x40\x59\x40\x9b\x46\x88\xf7\x09\x61\x70\x1e\x42\x54\x31\x85\x46\xfc\xf2\xac\xb4\x76\xc9\xc6\x4e\x5c\x9d\xbc\x05\xc9\x7b\x39\x5a\xbc\x1e\x5b\x38\x5e\xd0\xd3\x40\xd8\xe7\x73\x48\x61\x42\xdb\xcb\x85\xac\xd2\x91\x2e\x78\x7c\xfd\xd2\x34\x3b\xcc\x99\x34\x43\x06\x70\x43\xa5\xda\xe7\x5f\x11\xa7\x80\x3e\x13\x38\x7f\xaa\xd8\xf1\x4c\x01\xa2\x7a\x37\x08\xda\xd9\xa8\xc8\xf2\xf5\xcc\x49\xc6\x2c\xd0\x99\xa6\xc4\x85\x7b\x88\x38\x7e\x5a\xa8\x4a\x1a\x0d\x60\x5d\x04\xbc\x51\xe0\xd0\x25\xe7\xa5\x01\x36\x66\x98\x5f\xe0\xdb\xd7\xf2\x42\x3d\xec\x5a\x4a\xec\xc9\x40\x36\x19\xd3\xae\x92\x8f\x01\x72\xed\x56\x8d\x58\x3c\xf5\xb2\x06\xd1\x14\xf1\x16\x21\x59\xfa\x91\xf0\x19\x1a\x8e\x8a\xb8\xc1\xbe\xf7\x18\x42\x71\xcb\xcb\xd6\x3e\x0c\x39\xa9\x10\xae\xce\xf7\xc5\x53\x2f\x6b\xb4\x82\xf5\x04\x64\xe9\x70\xaf\xbb\x3f\x03\x88\xce\x40\xb3\x42\xde\x11\x9c\x45\x19\x59\xd6\xa7\x31\xee\x4a\xf6\xf4\x19\x58\x75\x2d\x54\xcd\xb5\xb0\x28\xae\x85\xaa\xb7\x82\x36\x43\x6d\x6c\x4f\xd2\xb5\x0e\xe8\x71\x50\xc9\x44\x10\x69\xd7\xaa\x4e\x5e\x5d\x54\x54\x1e\x92\x37\x45\xb4\x93\x77\x03\xb1\x64\x26\xd2\x31\x79\xe1\x9b\xcf\x5d\x8e\xf9\x7b\xa3\x85\x76\x6e\x13\x89\x3b\xf4\x70\xc8\x62\x38\x00\xc3\x1f\x58\xd7\x87\x12\x59\xc2\x76\xec\xb7\xbd\xd8\x56\x2f\xfb\xb1\x70\xd1\x88\xdc\xed\x38\xe5\x8f\x79\x3d\xa3\x95\xc5\x98\xe5\x7c\xe5\xb5\x92\x3d\x90\x89\xcf\xe7\x4e\xc5\xd5\xc0\x2a\x4a\xfc\xb9\x65\xd4\xc9\x7b\xcc\xf5\xd4\xa0\x85\x5c\x06\xa5\x55\x24\x67\x79\x33\x7b\x6e\xf7\x36\xa1\x66\x06\x88\x4e\xbc\x77\x59\xca\xb9\xa0\x92\xf8\x98\xd5\x95\x2c\xb1\xfd\xe5\x62\x44\xf2\xc2\x90\x0b\x90\xff\x0a\xb7\xb8\xe1\x96\x06\xf9\xd0\x3d\xa4\x3c\x6b\xf0\x1f\xbe\x61\x55\xe4\x43\xfd\xdb\xaf\xf8\x9f\x1c\x46\x05\xde\xe0\xa9\x5f\x86\x35\x9f\x3f\x34\xac\xb6\x11\x9d\x7a\x1a\x95\xbf\xc3\x77\xbc\xbf\x50\xff\xda\xf0\xff\xf3\x57\x00\x00\x00\xff\xff\x51\xc4\x11\xc1\x15\x06\x00\x00"

func migrations_postgres_0003_create_accounts_table_up_sql_bytes() ([]byte, error) {
	return bindata_read(
		_migrations_postgres_0003_create_accounts_table_up_sql,
		"migrations/postgres/0003_create-accounts-table_up.sql",
	)
}

func migrations_postgres_0003_create_accounts_table_up_sql() (*asset, error) {
	bytes, err := migrations_postgres_0003_create_accounts_table_up_sql_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "migrations/postgres/0003_create-accounts-table_up.sql", size: 1557, mode: os.FileMode(420), modTime: time.Unix(1425515484, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _migrations_postgres_0004_create_scopes_table_down_sql = "\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x4a\x29\xca\x2f\x50\x28\x49\x4c\xca\x49\x55\xc8\x4c\x53\x48\xad\xc8\x2c\x2e\x29\x56\x28\x4e\xce\x2f\x48\x2d\xb6\x06\x04\x00\x00\xff\xff\x5f\xc6\xb7\xd6\x1c\x00\x00\x00"

func migrations_postgres_0004_create_scopes_table_down_sql_bytes() ([]byte, error) {
	return bindata_read(
		_migrations_postgres_0004_create_scopes_table_down_sql,
		"migrations/postgres/0004_create-scopes-table_down.sql",
	)
}

func migrations_postgres_0004_create_scopes_table_down_sql() (*asset, error) {
	bytes, err := migrations_postgres_0004_create_scopes_table_down_sql_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "migrations/postgres/0004_create-scopes-table_down.sql", size: 28, mode: os.FileMode(420), modTime: time.Unix(1425515484, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _migrations_postgres_0004_create_scopes_table_up_sql = "\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x4f\x4b\x4e\xc4\x30\x0c\xdd\xe7\x14\xde\x31\x23\xcd\x70\x01\xb6\x6c\x38\xc5\xc8\x24\xae\x6a\x91\x26\x51\xfc\xaa\x99\x72\x7a\x4a\x5a\x02\x08\x24\xf0\xca\x4f\x7e\x3f\x9f\xcf\x84\x51\x8d\xc0\xcf\x51\xc8\xe7\x04\xd6\xb4\xc2\x51\x88\x63\x24\xf3\xb9\x88\x51\x1e\x36\x96\x2d\x06\x99\xee\x9d\xaf\xc2\x90\x5d\xa4\x03\xa5\x0c\x92\x9b\x1a\xec\x43\x71\x70\xb4\xce\xea\xde\x30\x3d\x3d\x36\xac\x81\x7e\x8c\x57\xc8\x0d\xcd\x22\xcd\x31\x9e\xbe\x09\xef\x8c\x82\x98\xaf\x5a\xa0\x39\xb5\xd3\x17\xdc\x3d\x7e\x77\x80\x4e\x62\xe0\xa9\xbc\xf7\xbf\x8e\x92\xda\x5b\x5b\xa1\x2b\x1b\x6d\x5f\x84\x46\xdf\xf7\x0b\xe3\xb3\x59\xd7\xe3\xb5\x9b\xaf\xf1\x03\xcf\x11\xe4\xe7\x5a\x25\xe1\xd2\x49\xff\x8e\x8d\x6c\xa0\xb9\x84\x9e\xbd\xef\x7f\x66\x9f\x5c\xa3\x97\xaa\x13\xd7\x85\x5e\x64\x39\x68\x38\xba\xe3\x83\x7b\x0b\x00\x00\xff\xff\x57\x58\x91\x50\xc6\x01\x00\x00"

func migrations_postgres_0004_create_scopes_table_up_sql_bytes() ([]byte, error) {
	return bindata_read(
		_migrations_postgres_0004_create_scopes_table_up_sql,
		"migrations/postgres/0004_create-scopes-table_up.sql",
	)
}

func migrations_postgres_0004_create_scopes_table_up_sql() (*asset, error) {
	bytes, err := migrations_postgres_0004_create_scopes_table_up_sql_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "migrations/postgres/0004_create-scopes-table_up.sql", size: 454, mode: os.FileMode(420), modTime: time.Unix(1425515484, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _migrations_postgres_0005_create_clients_table_down_sql = "\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x4a\x29\xca\x2f\x50\x28\x49\x4c\xca\x49\x55\xc8\x4c\x53\x48\xad\xc8\x2c\x2e\x29\x56\x48\xce\xc9\x4c\xcd\x2b\x29\xb6\x06\x04\x00\x00\xff\xff\xae\x13\x4f\x86\x1d\x00\x00\x00"

func migrations_postgres_0005_create_clients_table_down_sql_bytes() ([]byte, error) {
	return bindata_read(
		_migrations_postgres_0005_create_clients_table_down_sql,
		"migrations/postgres/0005_create-clients-table_down.sql",
	)
}

func migrations_postgres_0005_create_clients_table_down_sql() (*asset, error) {
	bytes, err := migrations_postgres_0005_create_clients_table_down_sql_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "migrations/postgres/0005_create-clients-table_down.sql", size: 29, mode: os.FileMode(420), modTime: time.Unix(1425515484, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _migrations_postgres_0005_create_clients_table_up_sql = "\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x92\xc1\x72\xdb\x40\x08\x86\xef\x7a\x0a\x6e\x6d\x66\x9c\x1c\x7a\xed\xb5\x97\x3e\x85\x87\xac\x90\xc5\x74\xb5\x68\x58\xb6\x8e\xf3\xf4\x65\xb5\x92\x23\x37\x99\xb1\x75\x12\x08\xbe\x1f\xc4\xff\xfc\x0c\x36\x72\x06\xc3\xd7\x48\x10\x24\x19\x72\xca\x80\x31\x82\x60\xb1\xf1\x07\x84\xc8\x94\x0c\x70\x9e\x23\x07\x34\x16\xff\xac\x74\xe2\x6c\xa4\xd4\x03\xa7\x06\xc8\x17\x4f\x4c\x2f\x5d\x50\x42\xa3\x95\xc7\x03\x24\x31\xa0\x37\xaf\xce\x2b\x29\xc3\xf7\x0e\xfc\x71\xe5\x15\xfd\xfb\xd7\xcb\x92\xe1\x1e\x3e\x3d\xa5\x78\xb6\x32\x52\x89\xf1\xb0\x35\x62\x08\x52\x96\x4e\x30\x81\xf3\xc8\x61\xf4\x29\x68\x03\xbe\x52\x94\x74\xf2\xa5\xa4\x81\xd7\xf2\xe3\x5e\xe0\x06\xec\x0b\x0d\xbe\x4d\x0a\x94\xb7\x62\x1f\x93\xfb\xa7\xc3\xed\xa8\xf9\x9b\x2f\x4a\xbe\xa2\x1d\xc0\xab\xf5\x32\x1b\xf5\x4d\xa3\xa5\x6f\x87\x37\x7a\xb3\xcf\xc3\x37\x94\x93\x12\x4e\xd4\x9a\xeb\xdb\xff\x9b\xdf\x69\xee\x29\x07\xe5\xb9\xde\xa3\x31\x76\x89\x47\x19\x7e\x53\x88\x72\x12\x28\x1a\x1b\xa4\x46\x47\x8f\xee\x0f\x32\xca\x44\x33\x9e\xa8\xf6\xc2\x20\xda\x5c\xd0\xd0\x8d\xb5\x55\xec\x79\x5f\xb3\xdc\x47\xac\x14\x6c\x61\x2d\x07\xf5\x63\x2c\x07\xad\x16\x14\xe5\xf7\xc5\x77\xfe\x8f\xf5\x2f\xe9\xb5\xbc\xfa\x30\x4b\xd1\x40\x20\xe7\x44\x9a\x9b\xee\xf6\x79\xaf\x1b\xf8\x6b\x65\xe3\x89\xb2\xe1\x34\x83\x0c\x55\x36\xed\x6d\x74\x46\x5f\x68\xb1\xf3\x7a\xe3\x35\x38\xe2\xee\xce\x57\x82\xbd\x7f\xd8\xa9\xa7\x01\x4b\x34\x08\x45\xdd\x55\x76\xbc\x16\x3d\x2e\x1c\x31\xfb\xff\x98\xfb\x0f\xf5\x35\xb8\xab\x7e\xe8\x96\xf2\x59\x79\x42\xbd\xc0\x1f\xba\x54\x27\x77\x4f\x3f\xbb\x7f\x01\x00\x00\xff\xff\xa6\xc0\x78\x00\xed\x03\x00\x00"

func migrations_postgres_0005_create_clients_table_up_sql_bytes() ([]byte, error) {
	return bindata_read(
		_migrations_postgres_0005_create_clients_table_up_sql,
		"migrations/postgres/0005_create-clients-table_up.sql",
	)
}

func migrations_postgres_0005_create_clients_table_up_sql() (*asset, error) {
	bytes, err := migrations_postgres_0005_create_clients_table_up_sql_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "migrations/postgres/0005_create-clients-table_up.sql", size: 1005, mode: os.FileMode(420), modTime: time.Unix(1425515484, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _migrations_postgres_0006_create_grants_table_down_sql = "\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x4a\x29\xca\x2f\x50\x28\x49\x4c\xca\x49\x55\xc8\x4c\x53\x48\xad\xc8\x2c\x2e\x29\x56\x48\x2f\x4a\xcc\x2b\x89\x2f\x4e\xce\x2f\x48\x2d\xb6\xe6\xc2\xad\x04\x2e\x59\x59\x80\x45\x7b\x49\x62\x49\x69\x71\x3c\x48\xce\x9a\x0b\x10\x00\x00\xff\xff\x3f\x5e\x37\x2d\x67\x00\x00\x00"

func migrations_postgres_0006_create_grants_table_down_sql_bytes() ([]byte, error) {
	return bindata_read(
		_migrations_postgres_0006_create_grants_table_down_sql,
		"migrations/postgres/0006_create-grants-table_down.sql",
	)
}

func migrations_postgres_0006_create_grants_table_down_sql() (*asset, error) {
	bytes, err := migrations_postgres_0006_create_grants_table_down_sql_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "migrations/postgres/0006_create-grants-table_down.sql", size: 103, mode: os.FileMode(420), modTime: time.Unix(1425515484, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _migrations_postgres_0006_create_grants_table_up_sql = "\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x54\x41\x8f\xe2\x3c\x0c\x3d\xd3\x5f\xe1\x1b\x20\x31\x73\xf8\x8e\xdf\x5e\x57\x5a\xad\xf6\x47\x54\x21\x75\x69\x34\x21\xa9\x12\x67\x0a\xf3\xeb\xd7\x4e\xd2\x32\xb0\xc3\x40\x25\x20\xa4\xf1\x7b\xcf\xf6\x73\x5e\x5e\x40\x07\x54\x84\x11\x94\x03\x74\xe9\x08\x74\x1e\x11\x7a\x1f\x20\x92\xa2\x14\x9b\xf2\xbe\x6c\x1f\x82\x72\xd4\x96\x17\x6d\xde\x51\xb1\x44\x6d\xd6\x0e\xa7\xf5\x0e\xd6\x29\x62\x27\xbf\x78\x1a\x4d\x28\xcb\x80\xef\xfe\x8d\x97\xdb\x1f\x4d\xc3\x84\x34\x98\x08\xa4\xf6\x16\x41\x7b\x47\xca\x38\x26\xb7\x16\xbc\x4a\x34\xfc\x57\x38\x40\xd6\x3e\x98\x0f\x45\xc6\x3b\x3e\xd7\xb1\xc2\x03\x3a\x0c\xac\xa5\xcb\xf2\x14\x68\x6b\xd0\x91\x40\xa6\x91\x0f\x05\x8c\x3e\x05\x8d\xe0\x27\x3e\x07\x31\x69\x8d\x31\xf6\xc9\x82\x1a\xc7\xe0\xdf\x95\x05\xf2\x42\x19\xd3\x11\x61\xe0\x23\x73\x44\x7c\x5d\xb2\xcc\xb2\x4c\x0f\xce\x13\xe0\xc9\x44\x8a\x45\x50\x84\x4d\xb3\x62\xa6\x5f\x59\x9d\xe8\x69\x56\xf2\x0d\x37\x4f\x4a\xa6\xcb\xc1\x2e\x59\xbb\xcb\x21\x4a\x6b\x9f\x38\xe8\xf7\x4f\xe1\x9f\x06\xa3\x07\xae\x01\x56\xf9\xb0\x47\xeb\xdd\x81\x2b\xe2\x5f\x19\x32\xef\xb5\x8c\x71\x07\x92\x35\xf7\x18\xd0\xb1\xe8\x0a\xc0\xca\x4c\xb7\x65\x2a\x60\xae\x5c\xf5\x52\x33\x32\xc7\xd2\x47\xe1\xba\x2e\xe7\x61\x49\x82\x29\x4b\xa3\x62\x6b\xdc\x85\x52\x62\xb9\xcb\xc7\x91\x3e\x6e\x92\x29\xbd\x07\xcd\x6e\xd9\x73\xa5\x1d\x7f\xfa\xcc\xd0\x7b\x6b\xfd\x64\xdc\xe1\x7f\x60\x27\xec\x40\x7c\xb0\x83\xea\x82\x1d\x54\x0f\x30\x5f\x45\xb8\x7a\xfe\xf5\xd5\x92\x6f\x87\xbd\x4a\x96\xa0\xf8\x2b\x6b\x60\x40\x06\xd5\x04\x29\xd8\x9c\xe1\xa5\xa6\x5f\xe5\x39\xb1\x47\x4b\x7b\x85\x7e\x0e\x6e\x25\xb8\x3e\xda\x10\x9e\xe8\x26\xd3\xa5\x06\x92\xe1\x34\xa0\xcb\x04\x5f\x41\xd6\x55\xab\xe8\xfb\x0a\x2e\xb9\xe8\x14\x82\xb4\x79\x39\xf4\x1c\xa3\x94\x14\xf6\xe7\x4f\xe6\x61\x72\xd9\xbc\x62\xfe\xb6\x7d\x8f\x28\x6a\xbf\x16\x5b\x3c\x4e\xea\x39\xdc\x4b\xfb\xeb\xea\x09\xdc\x66\x35\x06\x73\x54\xe1\x0c\x6f\x78\x86\x8d\xd8\x75\xdb\xd4\x3b\xe4\x0f\xe2\x18\x33\x87\xe5\x19\x15\xd2\xa8\xfd\x28\xb7\x58\x8c\x5e\x9b\x7c\x49\xf0\xb0\xb1\x4b\xef\x3b\xff\xd1\xc8\xb7\x15\x92\x07\xbf\xfc\xaf\xf3\x7e\x77\x1e\xe7\x8b\x22\x2b\xe5\xba\xe4\xf8\x79\x96\x6f\x3c\xf6\x39\x6e\xe6\xe1\x31\x96\xfc\xfe\x06\x00\x00\xff\xff\xfb\xea\x62\x48\x94\x05\x00\x00"

func migrations_postgres_0006_create_grants_table_up_sql_bytes() ([]byte, error) {
	return bindata_read(
		_migrations_postgres_0006_create_grants_table_up_sql,
		"migrations/postgres/0006_create-grants-table_up.sql",
	)
}

func migrations_postgres_0006_create_grants_table_up_sql() (*asset, error) {
	bytes, err := migrations_postgres_0006_create_grants_table_up_sql_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "migrations/postgres/0006_create-grants-table_up.sql", size: 1428, mode: os.FileMode(420), modTime: time.Unix(1425515484, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _migrations_postgres_0007_create_tokens_table_down_sql = "\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x4a\x29\xca\x2f\x50\x28\x49\x4c\xca\x49\x55\xc8\x4c\x53\x48\xad\xc8\x2c\x2e\x29\x56\x28\xc9\xcf\x4e\xcd\x2b\xb6\xe6\x82\x48\x56\x16\x60\xc8\xc5\x17\x97\x24\x96\x94\x16\xc7\x83\xe4\xb0\x2b\xcb\x4f\x2c\x2d\xc9\x88\x87\x28\x86\xa8\x02\x04\x00\x00\xff\xff\x9e\x2e\x75\xdb\x6a\x00\x00\x00"

func migrations_postgres_0007_create_tokens_table_down_sql_bytes() ([]byte, error) {
	return bindata_read(
		_migrations_postgres_0007_create_tokens_table_down_sql,
		"migrations/postgres/0007_create-tokens-table_down.sql",
	)
}

func migrations_postgres_0007_create_tokens_table_down_sql() (*asset, error) {
	bytes, err := migrations_postgres_0007_create_tokens_table_down_sql_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "migrations/postgres/0007_create-tokens-table_down.sql", size: 106, mode: os.FileMode(420), modTime: time.Unix(1425515484, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _migrations_postgres_0007_create_tokens_table_up_sql = "\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x54\xc1\x6e\xdb\x30\x0c\x3d\x37\x5f\x41\xf4\x92\x16\x70\x77\xd8\x71\x3b\x0d\xd8\xb0\xed\xb4\x4b\xef\x86\x2c\xd3\xb1\x50\x47\x32\x24\x2a\x6e\xf6\xf5\xa3\x48\xc5\x6d\x92\x6d\xad\x81\x00\x8a\x44\xbe\x47\xf1\x3d\xf1\xe1\x01\x6c\x44\x43\x98\xc0\x78\x40\x9f\xf7\x40\xc7\x19\x61\x08\x11\x28\x3c\xa1\x87\x44\x86\x72\xc2\xb4\xd1\x38\x3d\x96\xa3\x56\x8f\x5a\xd9\x31\x49\xb3\xef\xb6\x1e\x97\x6d\x03\x5b\x7c\x9e\x5d\xc4\xbe\x2c\x23\x1e\x38\xbe\xdf\xde\x7f\xde\x6c\xfe\x47\xf8\xeb\x4b\xa6\xf1\x63\xe5\x2d\xbb\xe7\xa4\xc1\xf0\x71\xab\xd4\x17\x9c\xc6\x5a\x4c\x49\xb9\x86\x88\x69\x3c\x71\xd1\xe8\x12\x90\xe9\x26\x04\x1b\x3c\x19\xe7\x99\x77\x9a\x14\xab\x52\x25\xd8\xa1\xc7\xc8\x3c\xbd\x94\x61\xc0\x4e\x0e\x3d\x71\x81\x7d\xc1\xd8\x45\xc3\x7f\x6c\xe8\xb1\x81\x3c\x07\x0f\x4c\x10\x72\xb4\x5c\xd1\xc2\x79\x60\xe6\x39\x86\x83\x99\x18\xad\x90\xa4\xbc\x47\x18\x79\xff\x14\x96\x60\x71\x34\x32\xec\xce\x1d\xd0\xaf\x88\xdb\x04\xc9\x86\x19\x3f\xac\x97\x94\x32\xdd\x00\x3e\x10\xe0\xb3\x4b\x94\x4e\x05\xde\x6d\x6e\x38\xed\x51\x1a\xc3\x4c\x99\x93\x6e\xb4\x4d\xe7\x5f\xce\xae\x97\x74\x9f\xa7\xa9\x79\x95\x54\xfa\xd5\xc8\xed\x7c\x58\x20\xf8\xe9\x08\xb7\x1d\x9a\x88\xf1\xb6\x22\x69\x4b\xd7\xcf\x3a\xc2\x67\xba\xc0\xba\x12\xa8\x01\xe4\xab\x95\x1e\x88\x00\x10\xca\xad\x45\x00\x0d\xe2\x32\xab\x6a\x67\xe0\x57\x4a\x9e\x68\xa0\xc7\xc1\xe4\x89\x60\x55\x54\x78\xab\x1e\x3f\xbf\x96\x16\x2f\xa3\xb3\x63\x15\x56\x0a\xe1\x45\x87\xce\xef\x5e\x64\x64\x5a\x4d\x69\xb9\x1d\xff\xe8\x4e\x29\x14\x23\xfa\xa2\x8f\x06\x73\x9b\x5d\x7f\xaf\x8c\xdf\x45\x73\xf6\x7d\x2f\xaa\x56\x81\xc6\xea\x7c\x7d\x1e\xfc\x6f\x70\x31\x11\x90\xdb\x17\x41\x44\xd5\xb6\xf8\xe4\x1d\x94\x12\xcc\x8c\x25\xbc\x72\x7e\xdb\x77\xd8\xf7\xc2\x28\x0f\x4f\xcc\x01\x8f\x67\x34\xe0\x88\xdd\xc4\x50\x9d\x96\x93\x0c\xef\x5d\xd8\xa9\x60\xfd\x08\x0b\x1e\x30\x36\x25\xde\x86\x3c\xf5\xdc\x1c\x82\x34\x46\xe7\x9f\xa0\x3b\x4a\x6e\xed\xaa\x58\x9a\x83\x3d\x65\xf1\x70\xe1\xae\x22\x62\xd2\xca\xba\xcc\x28\x18\xcb\xf3\x61\xe3\xa8\x3d\x67\x7e\x1d\x5c\xab\xf1\x47\x18\x72\x2c\x26\x60\x66\xa9\xe0\xc2\x93\x7f\x75\x92\x8e\x0e\xb0\x3c\x03\xf8\x26\xc1\xf3\x6f\xd0\x8e\x86\x69\x0a\x0b\xab\xf9\x09\x78\x90\x34\x50\xc7\x48\x03\x75\x88\x14\x12\xcd\x3d\xfb\xae\x07\xd2\xb5\xa5\x64\x30\x09\x7b\xe9\x24\xc7\xee\xe7\xc2\xba\x8c\xc5\xcc\xab\xb4\x0b\x4f\x15\xd5\x5b\x6c\xa4\xab\xd6\xd0\x2b\xae\x53\x36\xfd\xbe\x66\xb1\x39\xc6\x62\xbc\x35\xe8\x7d\x8c\x2f\xb7\xab\xab\x37\x19\x19\x17\x18\x58\xfa\x63\xc8\xb1\x86\xe2\x0f\x75\xe6\xe9\x6d\x30\x9e\x36\x30\xb5\xce\xbf\x85\xb7\xb9\x99\xa3\xdb\x9b\x78\x84\x27\x3c\xc2\x9d\x00\xdc\x6f\x78\x90\xfe\x09\x00\x00\xff\xff\xc0\xbc\xe5\xa9\x25\x06\x00\x00"

func migrations_postgres_0007_create_tokens_table_up_sql_bytes() ([]byte, error) {
	return bindata_read(
		_migrations_postgres_0007_create_tokens_table_up_sql,
		"migrations/postgres/0007_create-tokens-table_up.sql",
	)
}

func migrations_postgres_0007_create_tokens_table_up_sql() (*asset, error) {
	bytes, err := migrations_postgres_0007_create_tokens_table_up_sql_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "migrations/postgres/0007_create-tokens-table_up.sql", size: 1573, mode: os.FileMode(420), modTime: time.Unix(1425515484, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"migrations/migrations.go": migrations_migrations_go,
	"migrations/postgres/0001_create-index-function-if-not-exists_down.sql": migrations_postgres_0001_create_index_function_if_not_exists_down_sql,
	"migrations/postgres/0001_create-index-function-if-not-exists_up.sql": migrations_postgres_0001_create_index_function_if_not_exists_up_sql,
	"migrations/postgres/0002_create-extension-citext_down.sql": migrations_postgres_0002_create_extension_citext_down_sql,
	"migrations/postgres/0002_create-extension-citext_up.sql": migrations_postgres_0002_create_extension_citext_up_sql,
	"migrations/postgres/0003_create-accounts-table_down.sql": migrations_postgres_0003_create_accounts_table_down_sql,
	"migrations/postgres/0003_create-accounts-table_up.sql": migrations_postgres_0003_create_accounts_table_up_sql,
	"migrations/postgres/0004_create-scopes-table_down.sql": migrations_postgres_0004_create_scopes_table_down_sql,
	"migrations/postgres/0004_create-scopes-table_up.sql": migrations_postgres_0004_create_scopes_table_up_sql,
	"migrations/postgres/0005_create-clients-table_down.sql": migrations_postgres_0005_create_clients_table_down_sql,
	"migrations/postgres/0005_create-clients-table_up.sql": migrations_postgres_0005_create_clients_table_up_sql,
	"migrations/postgres/0006_create-grants-table_down.sql": migrations_postgres_0006_create_grants_table_down_sql,
	"migrations/postgres/0006_create-grants-table_up.sql": migrations_postgres_0006_create_grants_table_up_sql,
	"migrations/postgres/0007_create-tokens-table_down.sql": migrations_postgres_0007_create_tokens_table_down_sql,
	"migrations/postgres/0007_create-tokens-table_up.sql": migrations_postgres_0007_create_tokens_table_up_sql,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintree_t struct {
	Func func() (*asset, error)
	Children map[string]*_bintree_t
}
var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"migrations": &_bintree_t{nil, map[string]*_bintree_t{
		"migrations.go": &_bintree_t{migrations_migrations_go, map[string]*_bintree_t{
		}},
		"postgres": &_bintree_t{nil, map[string]*_bintree_t{
			"0001_create-index-function-if-not-exists_down.sql": &_bintree_t{migrations_postgres_0001_create_index_function_if_not_exists_down_sql, map[string]*_bintree_t{
			}},
			"0001_create-index-function-if-not-exists_up.sql": &_bintree_t{migrations_postgres_0001_create_index_function_if_not_exists_up_sql, map[string]*_bintree_t{
			}},
			"0002_create-extension-citext_down.sql": &_bintree_t{migrations_postgres_0002_create_extension_citext_down_sql, map[string]*_bintree_t{
			}},
			"0002_create-extension-citext_up.sql": &_bintree_t{migrations_postgres_0002_create_extension_citext_up_sql, map[string]*_bintree_t{
			}},
			"0003_create-accounts-table_down.sql": &_bintree_t{migrations_postgres_0003_create_accounts_table_down_sql, map[string]*_bintree_t{
			}},
			"0003_create-accounts-table_up.sql": &_bintree_t{migrations_postgres_0003_create_accounts_table_up_sql, map[string]*_bintree_t{
			}},
			"0004_create-scopes-table_down.sql": &_bintree_t{migrations_postgres_0004_create_scopes_table_down_sql, map[string]*_bintree_t{
			}},
			"0004_create-scopes-table_up.sql": &_bintree_t{migrations_postgres_0004_create_scopes_table_up_sql, map[string]*_bintree_t{
			}},
			"0005_create-clients-table_down.sql": &_bintree_t{migrations_postgres_0005_create_clients_table_down_sql, map[string]*_bintree_t{
			}},
			"0005_create-clients-table_up.sql": &_bintree_t{migrations_postgres_0005_create_clients_table_up_sql, map[string]*_bintree_t{
			}},
			"0006_create-grants-table_down.sql": &_bintree_t{migrations_postgres_0006_create_grants_table_down_sql, map[string]*_bintree_t{
			}},
			"0006_create-grants-table_up.sql": &_bintree_t{migrations_postgres_0006_create_grants_table_up_sql, map[string]*_bintree_t{
			}},
			"0007_create-tokens-table_down.sql": &_bintree_t{migrations_postgres_0007_create_tokens_table_down_sql, map[string]*_bintree_t{
			}},
			"0007_create-tokens-table_up.sql": &_bintree_t{migrations_postgres_0007_create_tokens_table_up_sql, map[string]*_bintree_t{
			}},
		}},
	}},
}}

// Restore an asset under the given directory
func RestoreAsset(dir, name string) error {
        data, err := Asset(name)
        if err != nil {
                return err
        }
        info, err := AssetInfo(name)
        if err != nil {
                return err
        }
        err = os.MkdirAll(_filePath(dir, path.Dir(name)), os.FileMode(0755))
        if err != nil {
                return err
        }
        err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
        if err != nil {
                return err
        }
        err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
        if err != nil {
                return err
        }
        return nil
}

// Restore assets under the given directory recursively
func RestoreAssets(dir, name string) error {
        children, err := AssetDir(name)
        if err != nil { // File
                return RestoreAsset(dir, name)
        } else { // Dir
                for _, child := range children {
                        err = RestoreAssets(dir, path.Join(name, child))
                        if err != nil {
                                return err
                        }
                }
        }
        return nil
}

func _filePath(dir, name string) string {
        cannonicalName := strings.Replace(name, "\\", "/", -1)
        return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

