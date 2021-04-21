  ifneq (,$(wildcard ../.env))
	include ../.env
	export
endif
MAINDIR = src
FRONTDIR = src/frontend
WEBSERVERDIR = src/webserver

.PHONY: node_modules test tidy

backend:
	cd $(MAINDIR) && go run main.go

frontend:
	# cd $(FRONTDIR) && npm run serve
	cd $(FRONTDIR) && npm run build
	cd $(WEBSERVERDIR) && go run server.go

# NOTE(Jovan): Should be deprecated?
dev: node_modules
	cd $(MAINDIR) && go run main.go &
	cd $(FRONTDIR) && npm run serve

install: node_modules

node_modules: $(FRONTDIR)/package.json
	cd $(FRONTDIR)
	npm install

test:
	cd $(MAINDIR) && go test -v ./...

tidy:
	cd $(MAINDIR) && go fmt ./...

# TODO(Jovan): Build npm
build:
	cd $(MAINDIR) && go build -v
