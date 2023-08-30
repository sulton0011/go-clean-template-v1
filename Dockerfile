FROM golang:1.19 as builder
#
RUN mkdir -p $GOPATH/src/gitlab.com/media_park/mb_go_order_service 
WORKDIR $GOPATH/src/gitlab.com/media_park/mb_go_order_service
# Copy the local package files to the container's workspace.
COPY . ./

# installing depends and build
RUN export CGO_ENABLED=0 && \
    export GOOS=linux && \
    go mod vendor && \
    make build && \
    mv ./bin/mb_go_order_service /

FROM alpine
COPY --from=builder mb_go_order_service .
RUN apk add --no-cache tzdata
ENV TZ Asia/Tashkent

CMD ["./mb_go_order_service"]