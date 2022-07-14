FROM golang:1.17

WORKDIR /app

RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct

RUN echo deb http://mirrors.tuna.tsinghua.edu.cn/ubuntu/ focal main restricted > /etc/apt/sources.list
RUN echo deb http://mirrors.tuna.tsinghua.edu.cn/ubuntu/ focal-updates main restricted >> /etc/apt/sources.list
RUN echo deb http://mirrors.tuna.tsinghua.edu.cn/ubuntu/ focal universe >> /etc/apt/sources.list
RUN echo deb http://mirrors.tuna.tsinghua.edu.cn/ubuntu/ focal-updates universe >> /etc/apt/sources.list
RUN echo deb http://mirrors.tuna.tsinghua.edu.cn/ubuntu/ focal multiverse >> /etc/apt/sources.list
RUN echo deb http://mirrors.tuna.tsinghua.edu.cn/ubuntu/ focal-updates multiverse >> /etc/apt/sources.list
RUN echo deb http://mirrors.tuna.tsinghua.edu.cn/ubuntu/ focal-backports main restricted universe multiverse >> /etc/apt/sources.list
RUN echo deb http://mirrors.tuna.tsinghua.edu.cn/ubuntu/ focal-security main restricted >> /etc/apt/sources.list
RUN echo deb http://mirrors.tuna.tsinghua.edu.cn/ubuntu/ focal-security universe >> /etc/apt/sources.list
RUN echo deb http://mirrors.tuna.tsinghua.edu.cn/ubuntu/ focal-security multiverse >> /etc/apt/sources.list

RUN gpg --keyserver keyserver.ubuntu.com --recv 3B4FE6ACC0B21F32
RUN gpg --export --armor 3B4FE6ACC0B21F32 | apt-key add -

RUN apt-get update -y
RUN apt-get -y install netcat

# 将当前所有文件复制到镜像内 /app 目录下
COPY . /app

EXPOSE 8080

RUN go build main.go dao.go cache.go

CMD ./main
