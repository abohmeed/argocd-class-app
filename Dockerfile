# Two stages, so the shipped image carries a single static binary and nothing else:
# no shell, no package manager, nothing to patch.
FROM golang:1.23-alpine AS build
WORKDIR /src
COPY go.mod ./
COPY *.go ./
# CGO off and a stripped binary, so it runs on scratch.
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o /storefront .

FROM scratch
COPY --from=build /storefront /storefront
EXPOSE 5678
ENTRYPOINT ["/storefront"]
