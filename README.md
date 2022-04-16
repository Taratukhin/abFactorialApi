#  HTTP Factorial service

## Task description

1. Please create public repository on github or any other public code repository website
2. Create a REST endpoint called calculate available at port 8989 so we can access it `http://localhost:8989/calculate`
3. Use [https://github.com/julienschmidt/httprouter](https://github.com/julienschmidt/httprouter) for creating a server 
4. Calculate endpoint has to take json with following structure `{"a":int,"b":int}` and calculate factorial of a and b using goroutines. [https://en.wikipedia.org/wiki/Factorial](https://en.wikipedia.org/wiki/Factorial)
5. Calculate will return json with the a! and b!
6. Please create middleware which will check if a and b exists in json and they are non-negative int in case of failure return json `{ "error":"Incorrect input"}` with error status code 400 Bad Request
7. Once the application is done please create a Dockerfile which runs the above code

## HOWTO

- Run `docker build -t server01 .`
- Run `docker run -d -p 8989:8989 server01`
- Run `curl -i -X GET -d "{\"a\":2,\"b\":10}" http://localhost:8989/calculate` Result must be `OK`
- Run `curl -i -X GET -d "{\"a\":3,\"b\":-10}" http://localhost:8989/calculate` Result must be `Bad Request`
- Run `curl -i -X GET -d "{\"a\":2}" http://localhost:8989/calculate` Result must be `Bad Request`

