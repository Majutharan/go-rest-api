package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Data declarations...

//post data...
var post1 = map[string] string{
	"userId": "1",
	"id": "1",
	"title": "sunt aut facere repellat provident occaecati excepturi optio reprehenderit",
	"body": "quia et suscipit\nsuscipit recusandae consequuntur expedita et cum\nreprehenderit molestiae ut ut quas totam\nnostrum rerum est autem sunt rem eveniet architecto",
}

var post2 = map[string] string{
	"userId": "1",
	"id": "2",
	"title": "qui est esse",
	"body": "est rerum tempore vitae\nsequi sint nihil reprehenderit dolor beatae ea dolores neque\nfugiat blanditiis voluptate porro vel nihil molestiae ut reiciendis\nqui aperiam non debitis possimus qui neque nisi nulla",
}

var post3 = map[string] string{
	"userId": "1",
	"id": "3",
	"title": "ea molestias quasi exercitationem repellat qui ipsa sit aut",
	"body": "et iusto sed quo iure\nvoluptatem occaecati omnis eligendi aut ad\nvoluptatem doloribus vel accusantium quis pariatur\nmolestiae porro eius odio et labore et velit aut",
}


// user data...
var user1 = map[string] string{
	"id": "1",
	"name": "sun",
	"username": "sun",
	"email": "sun@gmail.com",
}

var user2 = map[string] string{
	"id": "2",
	"name": "man",
	"username": "man",
	"email": "man@gmail.com",
}

// comment data...
var comment1 = map[string] string{
	"postId": "1",
	"id": "1",
	"name": "id labore ex et quam laborum",
	"email": "Eliseo@gardner.biz",
	"body": "laudantium enim quasi est quidem magnam voluptate ipsam eos\ntempora quo necessitatibus\ndolor quam autem quasi\nreiciendis et nam sapiente accusantium",
}

var comment2 = map[string] string{
	"postId": "1",
	"id": "2",
	"name": "quo vero reiciendis velit similique earum",
	"email": "Jayne_Kuhic@sydney.com",
	"body": "est natus enim nihil est dolore omnis voluptatem numquam\net omnis occaecati quod ullam at\nvoluptatem error expedita pariatur\nnihil sint nostrum voluptatem reiciendis et",
}

// create users list...
var UsersList = []map[string] string {
	user1, user2,
}

// create post list
var PostList = []map[string] string {
	post1,post2,post3,
}

// create comment list...
var CommentList = []map[string] string{
	comment1,comment2,
}



func main() {

	r := gin.Default()

	r.GET("/users", func(c *gin.Context) {
		c.JSON(http.StatusOK, UsersList)
	})

	r.GET("/posts", func(c *gin.Context) {
		c.JSON(http.StatusOK, PostList)
	})

	r.GET("/posts/:postId", func(c *gin.Context) {
		postId := c.Param("postId")
		var post map[string] string = nil
		for i := 0; i < len(PostList); i++ {
			if  PostList[i]["id"] == postId {
				post = PostList[i]
			}

		}
		c.JSON(http.StatusOK, post)
	})

	r.GET("/post", func(c *gin.Context) {
		userId := c.DefaultQuery("userId", "1")
		var post map[string] string = nil
		var postList [] map[string] string
		for i := 0; i < len(PostList); i++ {
			if  PostList[i]["userId"] == userId {
				post = PostList[i]
				postList = append(postList, post)
			}

		}
		c.JSON(http.StatusOK, postList)
	})

	r.GET("/posts/:postId/comments", func(c *gin.Context) {
		postId := c.Param("postId")
		var comment map[string] string = nil
		var commentList [] map[string] string
		for i := 0; i < len(CommentList); i++ {
			if  CommentList[i]["postId"] == postId {
				comment = CommentList[i]
				commentList = append(commentList, comment)
			}

		}
		c.JSON(http.StatusOK, commentList)
	})


	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
