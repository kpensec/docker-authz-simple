FROM golang:1.15-buster as builder
WORKDIR /src
COPY . .
RUN make docker-authz-plugin

FROM alpine
COPY --from=builder /src/docker-authz-plugin /bin/docker-authz-plugin

ENTRYPOINT [ "/bin/docker-authz-plugin" ]
