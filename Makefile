include .env

GO_BIN:=$(shell go env GOPATH)/bin
WD:=$(shell pwd)
MYSQL_INFO:=-h 127.0.0.1 -P 3306 -u root
DB_NAME:=sample
DML_DIR:=./migration/dml

.PHONY: help
help:
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.PHONY: run-db
run-db: ## create DB
	docker-compose up --build -d mysql

.PHONY: run-go
run-go: ## run app
	docker-compose up --build server

.PHONY: run-app
run-app: ## run app
	docker-compose up --build

.PHONY: down
down: ## down
	docker-compose down --rmi all --volumes --remove-orphans

# create DB table
.PHONY: migrate
migrate: ## migrate
	migrate -path migration/ddl/ -database 'mysql://root:@tcp(localhost:3306)/$(DB_NAME)?parseTime=true&loc=Local' up

# insert data
.PHONY: seed
seed: ## seed
	mysql $(MYSQL_INFO) $(DB_NAME) < $(DML_DIR)/dummy_students_insert.sql

# delete data
.PHONY: delete
delete: ## delete
	mysql $(MYSQL_INFO) $(DB_NAME) < $(DML_DIR)/dummy_students_delete.sql

# .PHONY: init-local-db
# init-local-db:
# 	$(MAKE) run-db
# 	$(MAKE) migrate
# 	$(MAKE) seed

# create model using sqlboiler
.PHONY: generate-model
generate-model:
	@$(GO_BIN)/sqlboiler mysql

# format code
.PHONY: fmt
fmt:
	gofmt -l -w .
