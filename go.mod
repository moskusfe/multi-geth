module github.com/ethereum/go-ethereum

go 1.12

require (
	github.com/Azure/azure-storage-blob-go v0.8.0
	github.com/Azure/go-autorest/autorest/adal v0.6.0 // indirect
	github.com/allegro/bigcache v0.0.0-20190218064605-e24eb225f156
	github.com/apilayer/freegeoip v0.0.0-20180702111401-3f942d1392f6
	github.com/aristanetworks/goarista v0.0.0-20170210015632-ea17b1a17847
	github.com/btcsuite/btcd v0.0.0-20171128150713-2e60448ffcc6
	github.com/cespare/cp v0.1.0
	github.com/davecgh/go-spew v1.1.1
	github.com/deckarep/golang-set v0.0.0-20180603214616-504e848d77ea
	github.com/docker/docker v0.0.0-20180625184442-8e610b2b55bf
	github.com/edsrzf/mmap-go v0.0.0-20160512033002-935e0e8a636c
	github.com/elastic/gosigar v0.0.0-20180330100440-37f05ff46ffa
	github.com/fatih/color v1.6.0
	github.com/fjl/memsize v0.0.0-20180418122429-ca190fb6ffbc
	github.com/gballet/go-libpcsclite v0.0.0-20190403181518-312b5175032f
	github.com/go-stack/stack v1.8.0
	github.com/golang/protobuf v1.3.1
	github.com/golang/snappy v0.0.1
	github.com/golangci/golangci-lint v1.18.0 // indirect
	github.com/gorilla/websocket v1.4.1
	github.com/graph-gophers/graphql-go v0.0.0-20190513003547-158e7b876106
	github.com/hashicorp/golang-lru v0.0.0-20160813221303-0a025b7e63ad
	github.com/howeyc/fsnotify v0.0.0-20151003194602-f0c08ee9c607 // indirect
	github.com/huin/goupnp v0.0.0-20161224104101-679507af18f3
	github.com/influxdata/influxdb v0.0.0-20180221223340-01288bdb0883
	github.com/jackpal/go-nat-pmp v0.0.0-20160603034137-1fa385a6f458
	github.com/julienschmidt/httprouter v0.0.0-20170430222011-975b5c4c7c21
	github.com/karalabe/usb v0.0.0-20190819132248-550797b1cad8
	github.com/kylelemons/godebug v1.1.0 // indirect
	github.com/mattn/go-colorable v0.1.0
	github.com/mattn/go-isatty v0.0.3
	github.com/mohae/deepcopy v0.0.0-20170929034955-c48cc78d4826
	github.com/naoina/go-stringutil v0.1.0 // indirect
	github.com/naoina/toml v0.0.0-20170918210437-9fafd6967416
	github.com/olekukonko/tablewriter v0.0.0-20170128050532-febf2d34b54a
	github.com/onsi/ginkgo v1.8.0 // indirect
	github.com/onsi/gomega v1.5.0 // indirect
	github.com/oschwald/maxminddb-golang v0.0.0-20180819230143-277d39ecb83e // indirect
	github.com/pborman/uuid v0.0.0-20170112150404-1b00554d8222
	github.com/peterh/liner v0.0.0-20190123174540-a2c9a5303de7
	github.com/prometheus/tsdb v0.0.0-20190402121629-4f204dcbc150
	github.com/rjeczalik/notify v0.9.1
	github.com/robertkrimen/otto v0.0.0-20170205013659-6a77b7cbc37d
	github.com/rs/cors v0.0.0-20160617231935-a62a804a8a00
	github.com/rs/xhandler v0.0.0-20160618193221-ed27b6fd6521 // indirect
	github.com/status-im/keycard-go v0.0.0-20190316090335-8537d3370df4
	github.com/steakknife/bloomfilter v0.0.0-20180922174646-6819c0d2a570
	github.com/steakknife/hamming v0.0.0-20180906055917-c99c65617cd3 // indirect
	github.com/stretchr/testify v1.3.0
	github.com/syndtr/goleveldb v0.0.0-20190318030020-c3a204f8e965
	github.com/tyler-smith/go-bip39 v0.0.0-20181017060643-dbb3b84ba2ef
	github.com/wsddn/go-ecdh v0.0.0-20161211032359-48726bab9208
	golang.org/x/crypto v0.0.0-20190313024323-a1f597ede03a
	golang.org/x/net v0.0.0-20190620200207-3b0461eec859
	golang.org/x/sync v0.0.0-20190423024810-112230192c58
	golang.org/x/sys v0.0.0-20190422165155-953cdadca894
	golang.org/x/text v0.3.0
	gopkg.in/check.v1 v1.0.0-20180628173108-788fd7840127
	gopkg.in/natefinch/npipe.v2 v2.0.0-20160621034901-c1b8fa8bdcce
	gopkg.in/olebedev/go-duktape.v3 v3.0.0-20180302121509-abf0ba0be5d5
	gopkg.in/sourcemap.v1 v1.0.5 // indirect
	gopkg.in/urfave/cli.v1 v1.20.0
	gopkg.in/yaml.v2 v2.2.2 // indirect
	gotest.tools v2.2.0+incompatible // indirect

)

// see https://github.com/golang/lint/issues/436#issuecomment-482066447
replace github.com/golang/lint v0.0.0-20190409202823-959b441ac422 => github.com/golang/lint v0.0.0-20190409202823-5614ed5bae6fb75893070bdc0996a68765fdd275
