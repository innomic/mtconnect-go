SHELL := /bin/bash

VERSION ?= "2.0"

.PHONY: clean
clean:
	echo 'clean'

.PHONY: codegen
codegen:
	bash schema/codegen.sh $(VERSION)
