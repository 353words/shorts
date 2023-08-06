## TODO

bp.go
caller.go
cfg.go
du.go
empty_select.go
expvar.go
file_srv.go
fwd.go
grades.go
gz.go
http_const.go
http_test.go
jscape.go
multierr.go
once.go
println.go
rate.go
report.go
req_ctx.go
rune_size.go
sha1.go
since.go
slice.go
srv_test.go
sql_named.go
tee.go
time_equal.go
user.go

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


## Done

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
