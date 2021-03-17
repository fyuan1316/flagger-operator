FROM harbor-b.alauda.cn/asm/builder:0.4-alpine3.12.1 AS builder

COPY ./bin/ /opt/

RUN ARCH="" && dpkgArch="$(arch)" \
  && case "${dpkgArch}" in \
  x86_64) ARCH='amd64' && upx /opt/${ARCH}/manager ;; \
  aarch64) ARCH='arm64' && upx /opt/${ARCH}/manager  ;; \
  *) echo "unsupported architecture"; exit 1 ;; \
  esac \
  && cp /opt/${ARCH}/manager /manager

# Use distroless as minimal base image to package the manager binary
# Refer to https://github.com/GoogleContainerTools/distroless for more details
#FROM gcr.io/distroless/static:nonroot
FROM alpine
RUN apk --no-cache --update add ca-certificates
WORKDIR /
COPY --from=builder /manager /manager
COPY files/ files/
#USER nonroot:nonroot
COPY --from=alpine/k8s:1.14.9 /usr/bin/kubectl /usr/local/bin/kubectl
RUN apk add --no-cache bash && rm -rf /var/cache/apk/*
ENTRYPOINT ["/manager"]
