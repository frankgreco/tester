ARG GO_VERSION=1.9.3
ARG ALPINE_VERSION=3.7

FROM golang:${GO_VERSION} AS BUILD
WORKDIR /go/src/github.com/frankgreco/tester/
COPY Gopkg.toml Gopkg.lock Makefile /go/src/github.com/frankgreco/tester/
RUN make install
COPY ./ /go/src/github.com/frankgreco/tester/
RUN CGO_ENABLED=0 make binary

FROM scratch
COPY --from=BUILD /go/src/github.com/frankgreco/tester/tester /
ENTRYPOINT ["/tester"]
