module github.com/scryinfo/go-substrate-rpc-client

go 1.14

require (
	github.com/aristanetworks/goarista v0.0.0-20190712234253-ed1100a1c015
	github.com/centrifuge/go-substrate-rpc-client v0.0.0-00010101000000-000000000000
	github.com/davecgh/go-spew v1.1.1
	github.com/deckarep/golang-set v1.7.1
	github.com/ethereum/go-ethereum v1.9.3
	github.com/go-stack/stack v1.8.0
	github.com/gorilla/websocket v1.4.1
	github.com/pierrec/xxHash v0.1.5
	github.com/pmezard/go-difflib v1.0.0
	github.com/rs/cors v1.6.0
	github.com/stretchr/testify v1.3.0
	golang.org/x/crypto v0.0.0-20190701094942-4def268fd1a4
	golang.org/x/sys v0.0.0-20190712062909-fae7ac547cb7
	gopkg.in/natefinch/npipe.v2 v2.0.0-20160621034901-c1b8fa8bdcce
)

replace github.com/centrifuge/go-substrate-rpc-client => ./
