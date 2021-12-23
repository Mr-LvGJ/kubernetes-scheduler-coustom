all: local

local: fmt vet
	GOOS=linux GOARCH=amd64 go build  -o sample-scheduler

build:  local
	docker build --no-cache . -t k765171999/sample-scheduler:v1.0.0

push:   build
	docker push registry.cn-hangzhou.aliyuncs.com/njupt-isl/yoda-scheduler:2.33

# Run go fmt against code
fmt:
	go fmt ./...

# Run go vet against code
vet:
	go vet ./...

clean: fmt vet
	sudo rm -f yoda-scheduler
