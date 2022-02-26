# go-rest-api-template

## How to run

`go run main.go` works fine if you have a single file or a script you're working on, but once you have a more complex project with lots of files then you'll have to start using the proper go build tool and run the compiled executable:

```bash
go build && ./api-service
```

Which is faster than first typing `go build` to generate an executable called `api-service` and then run that executable by typing `./api-service`.

The `api-service` app will bind itself to the port that you have defined in your environment variables or your Makefile (in our case `3001`). You can copy paste the following in your terminal to run the application (ensure you are in the `cmd/api-service` folder):

```bash
export ENV=LOCAL
export VERSION=VERSION
export PORT=3001
export FIXTURES=./fixtures.json
go build && ./api-service
```
