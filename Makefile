# 定义全局 Makefile 变量 方便后面引用
COMMON_SELF_DIR := $(dir $(lastword $(MAKEFILE_LIST)))

# 项目根目录
ROOT_DIR := $(abspath $(shell cd $(COMMON_SELF_DIR)/ && pwd -P))

# 构建产物、临时文件存放目录
OUTPUT_DIR := $(ROOT_DIR)/_output

## 指定应用使用的 version 包， 会通过 `-ldflgas -X` 向该包中指定的变量注入值
VERSION_PACKAGE=cn.xdmnb/study/miniblog/pkg/version

ifeq ($(origin VERSION) , undefined)
VERSION := $(shell git describe --tags --always --match='v*')
endif
## 检查代码仓库是否是 dirty（默认dirty）
GIT_TREE_STATE:="dirty"
ifeq (, $(shell git status --porcelain 2>/dev/null))
	GIT_TREE_STATE="clean"
endif
GIT_COMMIT:=$(shell git rev-parse HEAD)

GO_LDFLAGS += \
	-X $(VERSION_PACKAGE).GitVersion=$(VERSION) \
	-X $(VERSION_PACKAGE).GitCommit=$(GIT_COMMIT) \
	-X $(VERSION_PACKAGE).GitTreeState=$(GIT_TREE_STATE) \
	-X $(VERSION_PACKAGE).BuildDate=$(shell date -u +'%Y-%m-%dT%H:%M:%SZ')


# 定义 Makefile all 伪目标，执行`make` 时，默认会执行 all 伪目标
.PHONY: all
all: add-copyright format build

# 定义其他需要的 伪目标
.PHONY: build
# 编译源码，依赖 tidy 目标自动添加/移除依赖包
build: tidy 
	@go build -v -ldflags "$(GO_LDFLAGS)" -o $(OUTPUT_DIR)/miniblog $(ROOT_DIR)/cmd/miniblog/main.go

.PHONY: format 
format: 
	@gofmt -s -w ./

# 添加版权信息
.PHONY: add-copyright
add-copyright: 
	@addlicense -v -f $(ROOT_DIR)/scripts/boilerplate.txt $(ROOT_DIR) --skip-dirs=third_party,vendor,$(OUTPUT_DIR)

.PHONY: tidy
tidy: # 自动添加/移除依赖包.
	@go mod tidy

.PHONY: clean
clean: # 清理构建产物、临时文件等.
	@-rm -vrf $(OUTPUT_DIR)