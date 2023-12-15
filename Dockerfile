FROM --platform=$BUILDPLATFORM golang:1.21.5 as builder
ARG TARGETARCH
WORKDIR /app

COPY .  ./

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=${TARGETARCH} go build -o ./echo-server ./main.go

######## Start a new stage from scratch #######
FROM alpine:3.19.0

ENV HTTP_PORT=8080
ENV HTTPS_PORT=8443
ENV GRPC_PORT=50051
ENV GRPCS_PORT=50053

EXPOSE 8080 8443 50051 50053

RUN apk add --no-cache ca-certificates bash curl jq yq busybox-extras \
    && rm -rf /var/lib/apk/lists/*

USER 1001
SHELL ["/bin/bash"]

WORKDIR /app

COPY --from=builder /app/certs  /app/certs
COPY --from=builder /app/echo-server /app/echo-server

ENTRYPOINT ["/app/echo-server"]
