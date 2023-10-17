# 使用官方的Golang镜像作为基础镜像
FROM golang:1.20.7 AS builder

# 设置代理服务器地址和端口
ENV GOPROXY=https://goproxy.cn,direct

# 设置工作目录(转到)
WORKDIR /TinyTik

# 复制go.mod和go.sum以下载依赖
COPY go.mod .
COPY go.sum .

# 下载依赖
RUN go mod download

# 复制整个项目
COPY . .

# 构建可执行文件
RUN go build -o /app/main cmd/main.go

# 使用更小的镜像作为最终镜像
FROM ubuntu:latest

# 清除代理设置
# ENV GOPROXY=""

# 在 Ubuntu 镜像中添加时区
RUN apt-get update && \
    apt-get install -y tzdata && \
    rm -rf /var/lib/apt/lists/*

# 设置默认时区为 Asia/Shanghai
ENV TZ=Asia/Shanghai

# 创建符号链接以更改容器内的系统时区
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && \
    echo $TZ > /etc/timezone

# 安装 FFmpeg
RUN apt-get update && \
    apt-get install -y ffmpeg

# 设置工作目录
WORKDIR /TinyTik

# 复制编译好的可执行文件
COPY --from=builder /app/main /TinyTik/app/main

# 复制配置文件等（如果有的话）
COPY config/ ./config/

# 暴露应用程序需要的端口
EXPOSE 8081

# 启动应用程序
CMD ["/TinyTik/app/main"]
