## TODO

bp.go
caller.go
du.go
empty_select.go
expvar.go
file_srv.go
fwd.go
gz.go
http_const.go
http_test.go
jscape.go
multierr.go
once.go
println.go
report.go
req_ctx.go
rune_size.go
sha1.go
since.go
slice.go
srv_test.go
sql_named.go
tee.go
user.go
tabw.go
generics_ptr.go
report.go
enc.go
caller.go


## Done

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
