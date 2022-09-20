pkgdir = $(dir $(lastword $(MAKEFILE_LIST))) #获取列表最后一个文件所在目录
pkgname = $(lastword $(subst /, ,$(pkgdir))) #去掉目录的斜杠/

include test01/test.txt

$(info "$(MAKEFILE_LIST)") #打印MAKEFILE_LIST列表

include test01/001/test.txt

$(info "$(MAKEFILE_LIST)") #打印MAKEFILE_LIST列表

include test02/test.txt

$(info "$(MAKEFILE_LIST)") #打印MAKEFILE_LIST列表

include test02/002/test.txt

$(info "$(MAKEFILE_LIST)") #打印MAKEFILE_LIST列表

all:
	@echo "...all"
	@echo "all: $(pkgname)"

