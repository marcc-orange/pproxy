# A demo golang proxy

## Install

    $ go get github.com/marcc-orange/pproxy

## Run

    $ pproxy -target https://http2.golang.org/ -header-key test -header-value go

## Test

    $ curl -v http://localhost:8080/reqinfo

You should see `test: go` in the headers of the HTTP request.

## Help

    $ pproxy -h
    Usage of ./pproxy:
      -header-key string
            The key of the injected header
      -header-value string
            The value of the injected header
      -listen int
            The port to listen locally (default 8080)
      -target string
            The URL to proxy (default "https://http2.golang.org/")

## License

Public domain