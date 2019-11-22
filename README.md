# Retrieving content from a URL at intervals (example of using GoLang).
author: Adam Nielski

The app uses the following Go packages:

* Web framework: [echo](https://echo.labstack.com/)
* Configuration: [viper](https://github.com/spf13/viper)
* Dependency management: [dep](https://github.com/golang/dep)
* Testing: [testify](https://github.com/stretchr/testify)

## Project Structure

The download-url-in-go project has six main packages:

* `bindings`: The bindings directory will hold all of the protocol-specific data bindings for the application, such as form submission, query string, and JSON representations of application input
* `handlers`: The handlers package is where you store all your Echo web application-handler code and business logic for the application
* `models`: Within an application, you’ll have application-specific data structures that you will need to persist to a database. The models package will house the application-specific types.
* `renderings`: The renderings package will contain all the data structures and types that will be serialized back to the caller through the http.ResponseWriter
* `vendor`: contains all third party libraries

[Dependency inversion principle](https://en.wikipedia.org/wiki/Dependency_inversion_principle)
is followed to make these packages independent of each other and thus easier to test and maintain.

### Configuration

You Must have set up .env environment variables.

### Run

make dockerrun

###API

See all request in browser: [http://IP:8080/swagger/index.html](http://<ip>:8080/swagger/index.html)

check your IP "api" container
```
docker inspect --format='{{.Name}} - {{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' $(docker ps -aq)
```

### Try:

```
curl -si IP:8080/api/fetcher -X POST --header "Content-Type: application/json" -d '{"url":"http://httpbin.org/range/10","interval":1, "id":5}'  
```

```
curl -si IP:8080/api/fetcher/5/history -GET   
```

```
curl -si IP:8080/api/fetcher -GET   
```

```
curl -si IP:8080/api/fetcher/8 -X DELETE
``` 

### Testing

Tests are run in isolated mode, to run test with watcher just run command on your host
`bin/run test` - feature pending!!!    

**now:** make test



