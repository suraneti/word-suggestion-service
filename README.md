## Running

Make sure you have [Golang](https://golang.org/dl/)

```sh
create floder go/src at c:/Users/%USER_NAME%/
cd c:/Users/%USER_NAME%/go/src
git clone git@bitbucket.org:innocnx/word_suggestion.git
cd word_suggestion
go run main.go
```

## Endpoints

### Resource components
Resource components list 

| method    | resource                      | request body (json only)              | description                 |
|:----------|:------------------------------|:--------------------------------------|:----------------------------|
|`POST`     | `/suggestion`                 | {"word": "Some word"}                 | Return suggestion word      |