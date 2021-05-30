FROM golang:alpine AS build
WORKDIR /build
COPY . .
RUN addgroup microblog -g 1001 \
    && adduser microblog -u 1001 -G microblog -D
RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux \
    go build -ldflags "-s -w -extldflags '-static'" -o ./app
#RUN apk add upx \
#    && upx ./app

FROM scratch
COPY --from=build /etc/passwd /etc/passwd
COPY --from=build --chown=1001:1001 /build/app /app

ENTRYPOINT ["/app"]