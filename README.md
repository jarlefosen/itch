# ITCH

## The itch you don't need to scratch

Itch is a prebuilt docker image built `FROM SCRATCH` that contains ca-certificates from alpine and a user `nobody`.

Inspired by [The Go 1.11 web service Dockerfile](https://medium.com/@pierreprinetti/the-go-1-11-dockerfile-a3218319d191)

### How to use this

```dockerfile
# Do your stuff, build a static binary
FROM golang:alpine as builder
WORKDIR /go/src/app
COPY . .

RUN CGO_ENABLED=0 go build -o /app main.go

# Copy binary from builder to itch
FROM jarlefosen/itch
COPY --from=builder /app /app
ENTRYPOINT ["/app"]
```

See the example above in action

- **examples/helloworld/** says Hello World and tells which user is running.
- **examples/httpget/** Shows how ITCH differs from SCRATCH when requesting data over HTTPS.
- **examples/timezone/** Shows how ITCH differs from SCRATCH when loading location zones.

Test each of the examples with

```sh
$ make build # builds individual scratch and itch images
$ make run   # executes SCRATCH and ITCH versions to compare behaviour
$ make clean # attempts to remove docker images

# This produces dangling Docker images. Clean all with
$ alias dcleanup='([ -z "$(docker images -q -f dangling=true)" ] && echo "No dangling images") || (echo "Removing dangling images" && docker rmi $(docker images -q -f dangling=true))'
$ dcleanup
```
