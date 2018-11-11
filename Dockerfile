FROM alpine AS base

RUN mkdir /user && \
    echo 'nobody:x:65534:65534:nobody:/:' > /user/passwd && \
    echo 'nobody:x:65534:' > /user/group

RUN apk add --no-cache ca-certificates

FROM scratch AS final

COPY --from=base /user/group /user/passwd /etc/
COPY --from=base /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

USER nobody:nobody
