# shortener
URL shortener


## Building

`make`

## Running

`./urlshortener`

## Help

```
./Usage of ./urlshortener:
  -bind string
        Address to bind to, host:port (default "0.0.0.0:8000")
  -readt int
        Read timeout, seconds (default 30)
  -tpldir string
        Path to templates (default "templates")
  -urldb string
        Path to URLDB CSV file
  -writet int
        Write timeout, seconds (default 30)
```

Notes: urldb is not supported yet
