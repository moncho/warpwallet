FROM golang:1.15-alpine as builder
ARG GOOS=linux
WORKDIR /go/src/github.com/moncho/warpwallet/
COPY .    .
RUN GOOS=${GOOS} go build 

FROM scratch
COPY --from=builder /go/src/github.com/moncho/warpwallet/warpwallet .
CMD ["./warpwallet"]  
