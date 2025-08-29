FROM ubuntu:24.04

# 避免构建时交互提示
ENV DEBIAN_FRONTEND=noninteractive

# install python3 and git
RUN apt-get update && \
    apt-get install -y --no-install-recommends \
      python3 python3-pip python3-venv \
      ca-certificates curl git build-essential bash \
      pkg-config libseccomp-dev && \
    rm -rf /var/lib/apt/lists/*

# point the python command to python3
RUN update-alternatives --install /usr/bin/python python /usr/bin/python3 1

# install go
ARG GO_VERSION=1.24.3
RUN curl -fsSL https://go.dev/dl/go${GO_VERSION}.linux-amd64.tar.gz -o /tmp/go.tgz && \
    tar -C /usr/local -xzf /tmp/go.tgz && \
    rm /tmp/go.tgz

# setup Go
ENV GOROOT=/usr/local/go
ENV GOPATH=/go
ENV PATH=$PATH:/usr/local/go/bin:/go/bin

# 创建 GOPATH 目录
RUN mkdir -p /go/{pkg,bin,src}

WORKDIR /workspace

CMD ["sleep", "infinity"]