include .env

GO_BIN:=$(shell go env GOPATH)/bin
WD:=$(shell pwd)
MYSQL_INFO:=-h 127.0.0.1 -P 3306 -u root
DB_NAME:=sample
DML_DIR:=./migration/dml

.PHONY: help
help: ## 使い方
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.PHONY: run-db
run-db: ## DB作成
	docker-compose up --build -d mysql

.PHONY: run-go
run-go: ## 起動
	docker-compose up --build server

.PHONY: down
down: ## down
	docker-compose down

# テーブル作成
.PHONY: migrate
migrate: ## migrate
	migrate -path migration/ddl/ -database 'mysql://root:@tcp(localhost:3306)/$(DB_NAME)?parseTime=true&loc=Local' up

# ローカルデータ挿入のコマンド
.PHONY: seed
seed: ## seed
	mysql $(MYSQL_INFO) $(DB_NAME) < $(DML_DIR)/dummy_users_insert.sql

# ローカルデータDLETEのコマンド
.PHONY: delete
delete: ## delete
	mysql $(MYSQL_INFO) $(DB_NAME) < $(DML_DIR)/dummy_users_delete.sql

# sqlboilerでmodel自動生成
.PHONY: generate-model
generate-model:
	@$(GO_BIN)/sqlboiler mysql

# コードを整形
.PHONY: fmt
fmt:
	gofmt -l -w .

