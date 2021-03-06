PROJECT_NAME     := Pulumi Python SDK
LANGHOST_PKG     := github.com/pulumi/pulumi/sdk/python/cmd/pulumi-language-python
VERSION          := $(shell ../../scripts/get-py-version)

ifeq ($(PYTHON),)
	PYTHON := python
endif

ifeq ($(PIP),)
	PIP := pip
endif

PYENV := ./env
PYENVSRC := $(PYENV)/src

include ../../build/common.mk

ensure::
	pipenv --python 3 install --dev
	mkdir -p $(PYENVSRC)

build_package::
	rm -rf $(PYENVSRC) && cp -R ./lib/. $(PYENVSRC)/
	sed -i.bak "s/\$${VERSION}/$(VERSION)/g" $(PYENVSRC)/setup.py && rm $(PYENVSRC)/setup.py.bak
	if [ ! -z "$$(command -v pandoc)" ]; then \
		pandoc --from=markdown --to=rst --output="$(PYENVSRC)/README.rst" ../../README.md; \
	else \
		echo "warning: pandoc not found, generating empty README.rst"; \
		echo "" > "$(PYENVSRC)/README.rst"; \
	fi
	cd $(PYENVSRC) && pipenv run python setup.py build bdist_wheel --universal

build_plugin::
	go install -ldflags "-X github.com/pulumi/pulumi/sdk/python/pkg/version.Version=${VERSION}" ${LANGHOST_PKG}

build:: build_package build_plugin

lint::
	pipenv run pylint ./lib/pulumi --rcfile=.pylintrc
	golangci-lint run

install_package::
	cp ./cmd/pulumi-language-python-exec "$(PULUMI_BIN)"

install_plugin::
	GOBIN=$(PULUMI_BIN) go install \
		  -ldflags "-X github.com/pulumi/pulumi/sdk/python/pkg/version.Version=${VERSION}" ${LANGHOST_PKG}

install:: install_package install_plugin

test_fast::
	pipenv run pip install ./env/src
	pipenv run python -m unittest discover -s lib/test -v

dist::
	go install -ldflags "-X github.com/pulumi/pulumi/sdk/python/pkg/version.Version=${VERSION}" ${LANGHOST_PKG}
	cp ./cmd/pulumi-language-python-exec "$$(go env GOPATH)"/bin/
