MKDOCS   := squidfunk/mkdocs-material:9.6
ROOT_DIR := $(shell dirname $(shell dirname $(realpath $(lastword $(MAKEFILE_LIST)))))

build:
	find $(ROOT_DIR)/docs/.cache/plugin/social/ -maxdepth 1 -type f | xargs rm
	docker run --rm -t -v "$(ROOT_DIR):/mnt" -w /mnt/docs $(MKDOCS) -- build

watch:
	docker run --rm -t -v "$(ROOT_DIR):/mnt" -w /mnt/docs -p 8000:8000 $(MKDOCS) -- serve --watch-theme --dev-addr 0.0.0.0:8000

ci-build:
	pip install -r requirements.txt 
	mkdocs build
