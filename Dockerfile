# ---- Build stage (Go) ----
FROM golang:1.25.6 AS builder
WORKDIR /src

ARG USE_UPX=1
ARG UPX_VERSION=5.1.0

COPY src/go.mod src/go.sum ./
RUN go mod download
COPY src/ .

ENV CGO_ENABLED=0

RUN --mount=type=cache,target=/root/.cache/go-build \
    go build -trimpath -buildvcs=false -ldflags="-s -w" -o /out/matrix-rss .

# (optional) UPX
RUN if [ "$USE_UPX" = "1" ]; then set -eux; \
      apt-get update && apt-get install -y --no-install-recommends curl xz-utils ca-certificates && \
      arch="$(uname -m)"; case "$arch" in \
        x86_64|amd64) upx_arch=amd64 ;; aarch64|arm64) upx_arch=arm64 ;; \
        *) echo "Unsupported arch: $arch" >&2; exit 1 ;; esac; \
      curl -fsSL -o /tmp/upx.txz \
        "https://github.com/upx/upx/releases/download/v${UPX_VERSION}/upx-${UPX_VERSION}-${upx_arch}_linux.tar.xz"; \
      tar -xJf /tmp/upx.txz -C /tmp; \
      mv /tmp/upx-${UPX_VERSION}*/upx /usr/local/bin/upx; chmod +x /usr/local/bin/upx; \
      upx --lzma --best --no-progress /out/matrix-rss; upx -t /out/matrix-rss; \
      rm -rf /var/lib/apt/lists/* /tmp/upx* /tmp/upx-${UPX_VERSION}*; \
    fi

# ---- Runtime (Distroless) ----
FROM gcr.io/distroless/static

ARG MATRIX_RSS_VERSION
ARG IMAGE

# ---- metadata ----
LABEL \
  org.opencontainers.image.title="Matrix RSS Bot" \
  org.opencontainers.image.description="Matrix Bot for RSS feeds" \
  org.opencontainers.image.authors="Fingo2409" \
  org.opencontainers.image.maintainer="Fingo2409" \
  org.opencontainers.image.version="${MATRIX_RSS_VERSION}" \
  org.opencontainers.image.licenses="MIT" \
  org.opencontainers.image.url="${IMAGE}" \
  org.opencontainers.image.documentation="TBD" \
  org.opencontainers.image.source="https://github.com/Fingo2409/matrix-rss" \
  org.opencontainers.image.base.name="gcr.io/distroless/static"

WORKDIR /

# CA-Bundle for HTTPS
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

COPY --from=builder /out/matrix-rss /matrix-rss

ENTRYPOINT ["/matrix-rss"]
