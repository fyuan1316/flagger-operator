FROM harbor-b.alauda.cn/asm/builder:0.4-alpine3.12.1 AS builder

COPY ./bin/ /opt/

RUN ARCH="" && dpkgArch="$(arch)" \
  && case "${dpkgArch}" in \
  x86_64) ARCH='amd64' && upx /opt/${ARCH}/manager ;; \
  aarch64) ARCH='arm64' && upx /opt/${ARCH}/manager  ;; \
  *) echo "unsupported architecture"; exit 1 ;; \
  esac \
  && cp /opt/${ARCH}/manager /manager

FROM alpine
WORKDIR /
COPY --from=builder /manager /manager
COPY files/ files/

COPY --from=alpine/k8s:1.14.9 /usr/bin/kubectl /usr/local/bin/kubectl
RUN apk add --no-cache bash ca-certificates && rm -rf /var/cache/apk/*
ENTRYPOINT ["/manager"]
