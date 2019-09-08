package main

import (
	"encoding/json"
	"fmt"
	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-gonic/gin"
	"github.com/sudhabindu1/wtf1/models"
	"github.com/sudhabindu1/wtf1/modules"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"time"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func init()  {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	r := gin.Default()
	r.HTMLRender = ginview.Default()
	r.GET("/", indexHandler)
	r.GET("/message/:uid", indexHandlerWithUid)
	r.GET("/json/:uid", messageHandler)
	r.POST("/insert", insertMessage)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	_ = r.Run() // listen and serve on 0.0.0.0:8080
}


func indexHandler(c *gin.Context)  {
	m, err := modules.FindMessage()
	if err != nil {
		c.String(http.StatusNotFound, err.Error())
		return
	}

	permalink := fmt.Sprintf("/message/%s", m.Uid)

	c.HTML(http.StatusOK, "index.html", gin.H{
		"message": m.Message,
		"link": m.Link,
		"permalink": permalink,
		"color": m.Color,
	})
	return
}

func indexHandlerWithUid(c *gin.Context)  {
	uid := c.Param("uid")

	m, err := modules.FindMessageWithId(uid)
	if err != nil {
		c.String(http.StatusNotFound, err.Error())
		return
	}

	c.HTML(http.StatusOK, "index.html", gin.H{
		"message": m.Message,
		"link": m.Link,
		"color": m.Color,
	})
	return
}

func insertMessage(c *gin.Context)  {
	token := c.Request.Header.Get("token")
	if token != os.Getenv("AUTH_TOKEN") {
		c.String(http.StatusUnauthorized, "User is not authorized")
		return
	}
	if c.Request.Body != nil {
		b, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
		}

		m := models.RadioMessage{}
		err = json.Unmarshal(b, &m)
		if err != nil {
			c.String(http.StatusBadRequest, "uid should not be sent")
		}
		if m.Uid != "" {
			c.String(http.StatusBadRequest, "uid should not be sent")
		}
		m.Uid = RandStringRunes(10)
		err = modules.InsertMessage(&m)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
		}
		c.String(http.StatusOK, fmt.Sprintf("created. uid: %v", m.Uid))
	}
}


func messageHandler(c *gin.Context)  {
	uid := c.Param("uid")

	m, err := modules.FindMessageWithId(uid)
	if err != nil {
		c.String(http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, *m)
	return

}

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}