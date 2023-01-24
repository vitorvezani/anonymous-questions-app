# API Developments

## Libraries used
- [Gin Web Framework](https://github.com/gin-gonic)
- [SQLite](https://github.com/glebarez/sqlite)

# Setup and bootstrap

## Adding GIN to the project

`pwd` to show the current directory, if not

> PATH_TO_PROJECT/backend

if you opened anonymous-questions-app
`code backend/`

pwd should show `PATH_TO_PROJECT/backend`

1. create file `main.go` on root of the backend project
1. setup the `main` function which is the entry-point of your application

```
func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
```

### Run 

`go get -u github.com/gin-gonic/gin` in the terminal to download the gin dependency

when everything compile run:

`go run main.go`

access `requests.rest` and send the `ping` request

To see the final result `git checkout . && git checkout gin_added`