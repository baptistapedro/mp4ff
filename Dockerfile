FROM golang:1.19.1-buster as go-target
ADD . /mp4ff
WORKDIR /mp4ff/cmd/mp4ff-info
RUN go build

FROM golang:1.19.1-buster
COPY --from=go-target /mp4ff/cmd/mp4ff-info/mp4ff-info /
COPY --from=go-target /mp4ff/mp4/testdata/*.mp4 /testsuite/

ENTRYPOINT []
CMD ["/mp4ff-info", "@@"]
