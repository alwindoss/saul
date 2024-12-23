# https://github.com/GoogleContainerTools/distroless
FROM golang:1.23.4 AS builder
WORKDIR /go/src/app
COPY . .
RUN make setup
RUN make docker

FROM gcr.io/distroless/static-debian12
WORKDIR /root/
COPY --from=builder /go/src/app/bin/saul .
CMD [ "./saul" ]