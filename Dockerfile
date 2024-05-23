FROM golang:alpine AS builder

# 为我们的镜像设置必要的环境变量
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY=https://goproxy.cn,direct

# 移动到工作目录：/build
WORKDIR /build

# 将代码复制到容器中
COPY . .

# 在含go环境的镜像中将代码编译成二进制可执行文件 app
RUN go build -o main .

# 需要运行的命令
ENTRYPOINT ["/main", "-f", "etc/config.yaml"]