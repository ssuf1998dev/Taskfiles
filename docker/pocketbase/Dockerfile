FROM alpine:3.22.0 AS builder

ARG PB_VERSION=0.28.3
ARG PB_ARCH=amd64

RUN apk add --no-cache unzip ca-certificates

ADD https://github.com/pocketbase/pocketbase/releases/download/v${PB_VERSION}/pocketbase_${PB_VERSION}_linux_${PB_ARCH}.zip /tmp/pb.zip
RUN unzip /tmp/pb.zip -d /pb/

FROM alpine:3.22.0

COPY --from=builder /pb/ /pb/

EXPOSE 22258

CMD ["/pb/pocketbase", "serve", "--http=0.0.0.0:22258"]
