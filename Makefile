test:
	go test -v -tags postgres -cover ./...

generate:
	go generate

deps:
	go get github.com/lib/pq
	go get github.com/hooklift/assert
	go get github.com/jteeuwen/go-bindata

.PHONY: deps test generate
