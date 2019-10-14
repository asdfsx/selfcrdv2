# 使用 kubebuilder 创建CRD

## 创建项目

```
$ mkdir $GOPATH/src/selfcrdv2
$ cd $GOPATH/src/selfcrdv2
$ kubebuilder init --domain clustar.ai
go get sigs.k8s.io/controller-runtime@v0.2.2
go: finding sigs.k8s.io/controller-runtime v0.2.2
......
Running make...
make
go get sigs.k8s.io/controller-tools/cmd/controller-gen@v0.2.1
......
controller-gen object:headerFile=./hack/boilerplate.go.txt paths="./..."
go fmt ./...
go vet ./...
go build -o bin/manager main.go
Next: Define a resource with:
kubebuilder create api
```

* `kubebuilder`在当前目录下创建项目的框架，并下载相关的依赖
* `domain`参数定义了API的`domain`部分

## 创建API

```
$ kubebuilder create api --group "core" --version v1 --kind SelfCRDV2
Create Resource [y/n]
y
Create Controller [y/n]
y
Writing scaffold for you to edit...
api/v1/selfcrdv2_types.go
controllers/selfcrdv2_controller.go
Running make...
/Users/sunxia/gocode/bin/controller-gen object:headerFile=./hack/boilerplate.go.txt paths="./..."
go fmt ./...
go vet ./...
go build -o bin/manager main.go
```

* `kind`要求使用驼峰格式
* `group`与`domain`组合成完整的 group url
* `version`是API版本

## 本地运行controller

```
$ make install
$ make run ENABLE_WEBHOOKS=false
```

## 部署

```
$ make docker-build docker-push IMG=asdfsx/selfcrdv2:latest
$ make deploy IMG=asdfsx/selfcrdv2:latest
```