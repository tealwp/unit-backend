FROM golang

RUN apt-get update && apt-get install -y curl protobuf-compiler

# sqlc install
RUN curl -LO https://downloads.sqlc.dev/sqlc_1.23.0_linux_amd64.tar.gz
RUN tar -xvf sqlc_1.23.0_linux_amd64.tar.gz
RUN cp sqlc /go/bin
RUN rm -rf sqlc_1.23.0_linux_amd64.tar.gz 
RUN rm -rf sqlc


RUN go install -v golang.org/x/tools/gopls@latest
# proto compiler
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
RUN go install github.com/davidrjenni/reftools/cmd/fillstruct@latest

# zsh install
RUN sh -c "$(wget -O- https://github.com/deluan/zsh-in-docker/releases/download/v1.1.5/zsh-in-docker.sh)" -- \
    -t robbyrussell \
    -p git \
    -p ssh-agent \
    -p https://github.com/zsh-users/zsh-autosuggestions \
    -p https://github.com/zsh-users/zsh-completions
