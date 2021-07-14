package web

import (
	"errors"
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
	//r.Use(identification())

	read := r.Group("/api/v0")
	{
		read.GET("/posts", posts)
		read.GET("/posts/:id", post)

		read.GET("/tags/:id/posts", posts)

		read.GET("/tags", tags)
		read.GET("/tags/:id", tag)

	}

	write := r.Group("/api/v0")
	write.Use(authorization())
	{
		write.POST("/posts", createPost)
		write.PATCH("/posts/:id", updatePost)
		write.DELETE("/posts/:id", deletePost)

		write.POST("/tags", createTag)
		write.PATCH("/tags/:id", updateTag)
		write.DELETE("/tags/:id", deleteTag)

		write.POST("/users", createUser)
		write.PATCH("/users/:id", updateUser)
		write.DELETE("/users/:id", deleteUser)
	}

	// sensitivity read
	senRead := r.Group("/api/v0")
	{
		senRead.GET("/users", users)
		senRead.GET("/users/:id", user)
	}

	return r
}

// ********** middlewares **********

//func identification() gin.HandlerFunc {
//
//}

func authorization() gin.HandlerFunc {
	return func(context *gin.Context) {
	}
}

// ********** handlers **********

// posts handles both `/post` and `/tags/:id/posts`, distinguish by checking `id`
// /posts/?status=normal&order=published_at,desc;title,asc&page=1&pageSize=10
func posts(c *gin.Context) {
	scope := c.Query("scope")
	c1, err := scopeCondition(scope)
	if err != nil {
		c.JSON(http.StatusBadRequest, ParamError)
		return
	}

	tag := c.Param("id")
	c2, err := tagCondition(tag)
	if err != nil {
		c.JSON(http.StatusBadRequest, ParamError)
		return
	}

	orders := c.DefaultQuery("order", "published_at,desc")
	c3, err := orderCondition(orders)
	if err != nil {
		c.JSON(http.StatusBadRequest, ParamError)
		return
	}

	page := c.DefaultQuery("page", "1")
	pageSize := c.DefaultQuery("pageSize", "10")
	c4, err := paginationCondition(page, pageSize)
	if err != nil {
		c.JSON(http.StatusBadRequest, ParamError)
		return
	}

	cond := mergeCondition(c1, c2, c3, c4)
	var posts []repo.Post
	err = repo.FindAll(&posts, cond)
	if err != nil {
		log.Printf("DB error: %v\n", err)
	}

	c.JSON(200, posts)
}

func post(c *gin.Context) {
	id := c.Param("id")

	c2, err := tagCondition(tag)
	if err != nil {
		c.JSON(http.StatusBadRequest, ParamError)
		return
	}
	c.JSON(http.StatusForbidden, &Response{Code: "123", Msg: "123"})
}

func createPost(c *gin.Context) {
	title := c.PostForm("title")
	abstract := c.PostForm("abstract")
	content := c.PostForm("content")
	contentType := c.PostForm("contentType")
	toc := c.PostForm("toc")
	publishedAt := c.PostForm("publishedAt")
	status := c.PostForm("status")
	tags := c.PostForm("tags")
	viewPath := c.PostForm("viewPath")
	if title == "" || content == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": "001",
			"msg":  "title and content can't be empty",
		})
		c.Abort()
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
		p.Status = repo.PUBLISHED
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
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 002,
			"msg":  "DB error",
		})
		c.Abort()
	}
	c.JSON(http.StatusCreated, gin.H{
		"code": 000,
		"msg":  "created",
	})
}

func updatePost(c *gin.Context) {
	c.AbortWithStatus(http.StatusForbidden)
}

func deletePost(c *gin.Context) {
	c.AbortWithStatus(http.StatusForbidden)
}

// tag

func tags(c *gin.Context) {
	cond := repo.Condition{
		Query:  "1=1",
		Orders: []interface{}{"name asc"},
	}

	var tags []repo.Tag
	err := repo.FindAll(&tags, cond)
	if err != nil {
		log.Panicln("DB error")
	}

	c.JSON(200, tags)
}

func tag(c *gin.Context) {
	c.AbortWithStatus(http.StatusForbidden)
}

func createTag(c *gin.Context) {
	c.AbortWithStatus(http.StatusForbidden)
}

func updateTag(c *gin.Context) {
	c.AbortWithStatus(http.StatusForbidden)
}

func deleteTag(c *gin.Context) {
	c.AbortWithStatus(http.StatusForbidden)
}

// user

func users(c *gin.Context) {
	c.AbortWithStatus(http.StatusForbidden)
}

func user(c *gin.Context) {
	//cond := repo.Condition{
	//	Query: "",
	//}
	c.AbortWithStatus(http.StatusForbidden)
}

func createUser(c *gin.Context) {
	c.AbortWithStatus(http.StatusForbidden)
}

func updateUser(c *gin.Context) {
	c.AbortWithStatus(http.StatusForbidden)
}

func deleteUser(c *gin.Context) {
	c.AbortWithStatus(http.StatusForbidden)
}

// ********** utilities **********

type Response struct {
	Code string
	Msg  string
}

var ParamError = Response{
	Code: "PE",
	Msg:  "Param Error",
}

func mergeCondition(conds ...repo.Condition) repo.Condition {
	var query = ""
	var args, orders []interface{}
	var offset, limit int
	for _, c := range conds {
		if c.Query != "" {
			if query == "" {
				query = c.Query
			} else {
				query = query + " AND " + c.Query
			}
		}

		args = append(args, c.Args...)

		orders = append(orders, c.Orders...)

		if c.Offset != 0 {
			offset = c.Offset
		}
		if c.Limit != 0 {
			limit = c.Limit
		}
	}
	return repo.Condition{
		Query:  query,
		Args:   args,
		Orders: orders,
		Offset: offset,
		Limit:  limit,
	}
}

func scopeCondition(scope string) (c repo.Condition, e error) {
	switch strings.ToLower(scope) {
	case "":
		fallthrough
	case "normal":
		c = repo.Condition{
			Query: "status & ? = ?",
			Args:  []interface{}{repo.PUBLISHED, repo.PUBLISHED},
		}
		e = nil
	case "all":
		c = repo.Condition{}
		e = nil
	default:
		c = repo.Condition{}
		e = errors.New("param error")
	}
	return
}

func tagCondition(tag string) (c repo.Condition, e error) {
	if tag == "" {
		c = repo.Condition{}
		e = nil
		return
	}

	tid, err := strconv.Atoi(tag)
	if err != nil {
		c = repo.Condition{}
		e = errors.New("tag id not number")
		return
	}

	c = repo.Condition{
		Query: "tags & ? = ?",
		Args:  []interface{}{tid, tid},
	}
	return
}

func orderCondition(orders string) (c repo.Condition, e error) {
	var ods []interface{}
	for _, o := range strings.Split(orders, ";") {
		ods = append(ods, strings.ReplaceAll(o, ",", " "))
	}
	c = repo.Condition{
		Orders: ods,
	}
	e = nil
	return
}

func paginationCondition(page, pageSize string) (c repo.Condition, e error) {
	p, err := strconv.Atoi(page)
	if err != nil {
		c = repo.Condition{}
		e = errors.New("page is not number")
		return
	}

	ps, err := strconv.Atoi(pageSize)
	if err != nil {
		c = repo.Condition{}
		e = errors.New("pageSize is not number")
		return
	}

	offset, limit := paginate(p, ps)
	c = repo.Condition{
		Offset: offset,
		Limit:  limit,
	}
	e = nil
	return
}

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
