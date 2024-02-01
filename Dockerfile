# Build binary
FROM golang:alpine as build

WORKDIR /build
# Use this if you have your go.mod in this folder instead
COPY go.mod go.sum ./
COPY . .
ENV GO111MODULE=on
ENV GOPATH=""
# RUN go mod init example.com/goLangMicroservice
# RUN go get -u github.com/gofiber/fiber/v2
RUN GOOS=linux go build -o ./main main.go

# Start lambda container from fresh image 
FROM golang:alpine


WORKDIR "/var/task"
COPY --from=build /build/main /var/task

COPY --from=public.ecr.aws/awsguru/aws-lambda-adapter:0.7.0 /lambda-adapter /opt/extensions/lambda-adapter




CMD ["/var/task/main"]
