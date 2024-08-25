module example.com/go_rest_api_test

go 1.18

require (
	example.com/helloworld v0.0.0-00010101000000-000000000000
	github.com/mattn/go-sqlite3 v1.12.0
	github.com/gorilla/mux v1.8.1
)

replace example.com/helloworld => ./helloworld

replace github.com/mattn/go-sqlite3 => ./vendor/github.com/mattn/go-sqlite3

replace github.com/gorilla/mux => ./vendor/github.com/gorilla/mux
