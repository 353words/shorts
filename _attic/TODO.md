## TODO

empty_select.go
enc.go
expvar.go
fin.go
fwd.go
generics_ptr.go
jscape.go
multierr.go
once.go
println.go
report.go
since.go
srv_test.go
tabw.go
tee.go
user.go


## Done

## Fourth Recording

http_test
    run & curl http://localhost:8080/api/health

num_chars
    https://go.dev/blog/strings

breakpoint
    Run with debugger
    Run without debugger

embed_fs
    http://localhost:8080/static/logo.svg

full_slice
    https://go.dev/ref/spec#Slice_expressions

sha1
    https://pkg.go.dev/crypto/sha1
    sha1sum httpd.log.gz

sql_named
    Not all drivers implement it.

http_const
    https://pkg.go.dev/net/http#pkg-constants

caller
    https://github.com/pkg/errors

zero
    https://github.com/golang/go/issues/61372
    
gz
    curl http://localhost:8080/agents
    curl -H 'Accept-Encoding: gzip' http://localhost:8080/agents
    curl -H 'Accept-Encoding: gzip' http://localhost:8080/agents | gunzip

### Third Recording


config_dir

    viper
    ardanlabs/conf

equal_time

    monotonic clock

greedy

    MustCompile

    https://regex101.com/
    https://pkg.go.dev/regexp/syntax

race

    $ go test -race counter_test.go

    https://pkg.go.dev/sync#Mutex
    https://pkg.go.dev/sync/atomic#AddInt64

    https://go.dev/doc/articles/race_detector#Runtime_Overheads
    use in tests

rate

    https://pkg.go.dev/golang.org/x/time/rate#Limiter
    https://github.com/rakyll/hey

    hey -n 1000 http://localhost:8080 

rename_import

    If use url can't call variable url
    Don't do this in general

request_ctx

    https://pkg.go.dev/net/http#Request.Clone
    curl http://localhost:8080

    https://github.com/ardanlabs/service/blob/master/foundation/web/context.go

search

    https://pkg.go.dev/sort#Search

skip

    $ go test -v
    $ CI=yes go test -v


### Second Recording

client

discard
    
    Discard HTTP request body
    https://pkg.go.dev/io#Copy

for_range

fmt_num

json_stream

    $ go run .

    https://jsonlines.org/

    Add

    fmt.Println(net.String())

    https://en.wikipedia.org/wiki/Chunked_transfer_encoding
    ref flusher short

limit

    $ cat /dev/random| head -c 200 > /tmp/req
    $ curl -i --data-binary @/tmp/req http://localhost:8080 
    ...
    read 200 bytes
    $ cat /dev/random| head -c 2000000 > /tmp/req
    $ curl -i --data-binary @/tmp/req http://localhost:8080
    ...
    can't read

    https://pkg.go.dev/io#LimitedReader

markdown_list
    https://pkg.go.dev/text/template
    https://pkg.go.dev/html/template

nfc

    https://en.wikipedia.org/wiki/Unicode_equivalence

trimpath

tz

    https://infiniteundo.com/post/25326999628/falsehoods-programmers-believe-about-time
    https://www.zainrizvi.io/blog/falsehoods-programmers-believe-about-time-zones/

word_freq



### First Recording

io_multi
json_binary
http_flusher
fmt_star
fmt_loc
stringer
regexp_fn
net_ip
compound_key
fold
is_alive
raw_string
pipe
label
cleanup
server
test_main
url_query
make
expandenv
