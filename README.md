# `go-probot`
> :construction:

## Example

Check out [the basic example](example/basic.go).

## Development

Use [`reflex`](https://github.com/cespare/reflex) to auto-reload the webhook server when any file changes are detected:

```shell
reflex -s -- go run cmd/go-probot/main.go -p 8888
```

Send an example webhook:

```shell
curl -X POST -d '{"hello":"world"}' -H "Content-type: application/json" http://localhost:8888/
```

Expose the endpoint over the public internet using `ngrok`:

```shell
ngrok http 8000
```

Alternatively, setup an [`ngrok` configuration file](https://ngrok.com/docs#config):

```yaml
tunnels:
  main:
    proto: http
    addr: 8888
```

And then:

```shell
ngrok start main
```
