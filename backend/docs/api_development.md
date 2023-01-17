# API Developments

## Libraries used
- [Gin Web Framework](https://github.com/gin-gonic)
- [SQLite](https://github.com/glebarez/sqlite)

## Installation

`go get -u github.com/gin-gonic/gin`

## Step-by-step

Each phase of the tutorial will be represented on a branch, you can move through the phases by changing the branch

`git checkout adding-server`

1. create file `main.go` on root of the backend project
2. setup the main function which is the entry-point of your application

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

To see the final result `git checkout adding-server-final`