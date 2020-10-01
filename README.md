## Running

Make sure you have [Golang](https://golang.org/dl/)

```sh
[Windows]
create floder go/src at c:/Users/%USER_NAME%/
cd c:/Users/%USER_NAME%/go/src
git clone https://github.com/suraneti/word-suggestion-service.git
cd word-suggestion-service
go run main.go

[MacOS]
cd /Users/$USER/go/src
git clone https://github.com/suraneti/word-suggestion-service.git
cd word-suggestion-service
go run main.go
```

## Build for Linux

### Git bash
```
GOOS=linux GOARCH=amd64 go build
```

## Endpoints

### Resource components
Resource components list 

| method    | resource                      | request body (json only)              | description                 |
|:----------|:------------------------------|:--------------------------------------|:----------------------------|
|`POST`     | `/suggestion`                 | {"word": "Some word"}                 | Return suggestion word      |
