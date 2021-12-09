# Country Rest API


🌍 *[English](README.md)

[![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/rs/rest-layer) 
[![Swagger Doc](https://godoc.org/github.com/swaggo/swagg?status.svg)](https://godoc.org/github.com/swaggo/swag)

Rest API was developed using standard libraries. Data is kept inMemory. 

HTTP requests are saved in the logFile.log file.

The data stored in the memory is saved to the countries.json file every 10 minutes.

```shell
    var interval = 10 * time.Minute
```


## How to run the application

1. Clone the application with `git@github.com/GamzeUnal/CountryRestAPI.git`

2. Download the postman configuration file (optional).
3. To build from source you need Go (1.13 or newer).
4. There are number of dependencies which need to be imported before running the application. Please get the dependenices through the following commands -

    ```shell
        go get -u "github.com/swaggo/swag/cmd/swag"
        go install "github.com/swaggo/swag/cmd/swag@latest"
    ```
5. To run the application, please use the following command -

    ```shell
        swag init
        go build
        go run
    ```
> Note: By default the port number its being run on is **5003**.

## Endpoints Description

### Get All Countries

```
    URL - *http://localhost:5003/countries*
    Method - GET
```

### Get Country By Code

```JSON
    URL - *http://localhost:5003/countries/?Code=tr*
    Method - GET
```

### Create Country

```JSON
    URL - *http://localhost:5003/newcountry/*
    Method - POST
    Body - (content-type = application/json)
    {	"Code" : "tr",
	"Name": "Türkiye", 
	"ParentCity": "Ankara"
	}
```

### Delete All Country(In Memory)

```JSON
    URL - *http://localhost:5003/countries/delete*
    Method - DELETE
```


## Test Driven Development Description

To run all the unit test cases, please do the following -

1. `go build`
2. `go test`


## Hope everything works. Thank you.
