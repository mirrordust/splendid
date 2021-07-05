package web

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"

	"github.com/mirrordust/splendour/m0/repo"
)

func Server() *gin.Engine {
	r := gin.Default()
	v0 := r.Group("/api/v0")

	v0.GET("/posts", posts)
	v0.GET("/posts/:id", post)
	v0.POST("/posts", createPost)
	v0.PATCH("/posts/:id", updatePost)
	v0.DELETE("/posts/:id", deletePost)

	v0.GET("/tags/:id/posts", posts)

	v0.GET("/tags", tags)
	v0.GET("/tags/:id", tag)
	v0.POST("/tags", createTag)
	v0.PATCH("/tags/:id", updateTag)
	v0.DELETE("/tags/:id", deleteTag)

	v0.GET("/users", users)
	v0.GET("/users/:id", user)
	v0.POST("/users", createUser)
	v0.PATCH("/users/:id", updateUser)
	v0.DELETE("/users/:id", deleteUser)

	return r
}

// handlers
// post

func posts(context *gin.Context) {
	// basic query and args
	status := context.DefaultQuery("status", "normal")
	s := repo.Status2Code(status)
	query := "status & ? = ?"
	args := []interface{}{s, s}
	// posts for a particular tag
	tag := context.Param("id")
	if tag != "" {
		tid, err := strconv.Atoi(tag)
		if err != nil {
			log.Println("Atoi tag id error")
		} else {
			query = query + " AND tags & ? = ?"
			args = append(args, tid, tid)
		}
	}

	// order
	orders := context.DefaultQuery("order", "published_at,desc")
	var ods []interface{}
	for _, o := range strings.Split(orders, ";") {
		ods = append(ods, strings.ReplaceAll(o, ",", " "))
	}

	// page
	page := context.DefaultQuery("page", "1")
	pageSize := context.DefaultQuery("pageSize", "10")
	pg, err := strconv.Atoi(page)
	if err != nil {
		log.Println("Atoi page error")
		pg = 1
	}
	pgs, err := strconv.Atoi(pageSize)
	if err != nil {
		log.Println("Atoi pageSize error")
		pgs = 10
	}
	offset, limit := paginate(pg, pgs)

	// query condition
	cond := repo.Condition{
		Query:  query,
		Args:   args,
		Orders: ods,
		Offset: offset,
		Limit:  limit,
	}
	var posts []repo.Post
	err = repo.FindAll(&posts, cond)
	if err != nil {
		log.Panicln("DB error")
	}

	context.JSON(200, posts)
}

func post(context *gin.Context) {
	context.AbortWithStatus(http.StatusForbidden)
}

func createPost(context *gin.Context) {
	title := context.PostForm("title")
	abstract := context.PostForm("abstract")
	content := context.PostForm("content")
	contentType := context.PostForm("contentType")
	toc := context.PostForm("toc")
	publishedAt := context.PostForm("publishedAt")
	status := context.PostForm("status")
	tags := context.PostForm("tags")
	viewPath := context.PostForm("viewPath")
	if title == "" || content == "" {
		context.JSON(http.StatusBadRequest, gin.H{
			"code": "001",
			"msg":  "title and content can't be empty",
		})
		context.Abort()
	}
	if viewPath == "" {
		viewPath = uuid.NewV4().String()
	}
	p := repo.Post{
		Title:       title,
		Abstract:    abstract,
		Content:     content,
		ContentType: contentType,
		TOC:         toc,
	}

	if publishedAt != "" {
		pa, err := strconv.ParseInt(publishedAt, 10, 64)
		if err != nil {
			log.Panicln("ParseInt publishedAt error")
		}
		p.PublishedAt = pa
	}

	if status != "" {
		s, err := strconv.ParseUint(status, 10, 8)
		if err != nil {
			log.Panicln("ParseInt status error")
		}
		p.Status = uint8(s)
	} else {
		p.Status = repo.NORMAL
	}

	if tags != "" {
		t, err := strconv.ParseUint(tags, 10, 64)
		if err != nil {
			log.Panicln("ParseInt tags error")
		}
		p.Tags = t
	}

	p.ViewPath = viewPath

	err := repo.Create(&p)
	if err != nil {
		log.Println(err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{
			"code": 002,
			"msg":  "DB error",
		})
		context.Abort()
	}
	context.JSON(http.StatusCreated, gin.H{
		"code": 000,
		"msg":  "created",
	})
}

func updatePost(context *gin.Context) {
	context.AbortWithStatus(http.StatusForbidden)
}

func deletePost(context *gin.Context) {
	context.AbortWithStatus(http.StatusForbidden)
}

// tag

func tags(context *gin.Context) {
	cond := repo.Condition{
		Query:  "1=1",
		Orders: []interface{}{"name asc"},
	}

	var tags []repo.Tag
	err := repo.FindAll(&tags, cond)
	if err != nil {
		log.Panicln("DB error")
	}

	context.JSON(200, tags)
}

func tag(context *gin.Context) {
	context.AbortWithStatus(http.StatusForbidden)
}

func createTag(context *gin.Context) {
	context.AbortWithStatus(http.StatusForbidden)
}

func updateTag(context *gin.Context) {
	context.AbortWithStatus(http.StatusForbidden)
}

func deleteTag(context *gin.Context) {
	context.AbortWithStatus(http.StatusForbidden)
}

// user

func users(context *gin.Context) {
	context.AbortWithStatus(http.StatusForbidden)
}

func user(context *gin.Context) {
	//cond := repo.Condition{
	//	Query: "",
	//}
	context.AbortWithStatus(http.StatusForbidden)
}

func createUser(context *gin.Context) {
	context.AbortWithStatus(http.StatusForbidden)
}

func updateUser(context *gin.Context) {
	context.AbortWithStatus(http.StatusForbidden)
}

func deleteUser(context *gin.Context) {
	context.AbortWithStatus(http.StatusForbidden)
}

// util functions

func paginate(page, pageSize int) (offset, limit int) {
	if page <= 0 {
		page = 1
	}
	switch {
	case pageSize > 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}

	return (page - 1) * pageSize, pageSize
}
