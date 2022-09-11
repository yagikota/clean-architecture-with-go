# clean_architecture_with_go
## 🗄 About this repository
* This is a sample API buil by Go(echo)and SQLBoiler according to Clean Architecture.

## 👟 How to run
### install
```shell
git clone git@github.com:yagikota/clean_architecture_with_go.git

cd clean_architecture_with_go
```
### set environment variables
```shell
echo "PORT=8080

MYSQL_USER=root
MYSQL_ALLOW_EMPTY_PASSWORD=true
MYSQL_PORT=3306
MYSQL_ADDR=mysql:3306
MYSQL_DATABASE=sample" > .env
```

### initialize local DB
```shell
make run-db
make migrate 
make seed
```
`make migrate` may fail, but try some time.

### build API
```shell
make run-go
```
then, access`localhost:8080/health_check`.
you will get
```json
{
    "message": "Hello! you've requested: /health_check"
}
```
🐶 you have fully prepared for running API

## 📄 API Document
* copy and paste [this file](https://github.com/yagikota/clean_architecture_with_go/blob/main/api_doc.yml) on https://editor.swagger.io/
