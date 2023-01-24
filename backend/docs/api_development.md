# API Developments

## Libraries used
- [Gin Web Framework](https://github.com/gin-gonic)
- [SQLite](https://github.com/glebarez/sqlite)

# Initial checks

- `pwd` should show `PATH_TO_PROJECT/backend`


```golang
	_, err := gorm.Open(sqlite.Open("sqlite.db"), &gorm.Config{})
	if err != nil {
		logrus.Fatal("could not open db connection", err)
	}
```

# Setup and bootstrap

## Adding SQLite and GORM to the project

1. create a `pkg` folder in backend

## Adding GIN to the project


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