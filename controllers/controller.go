package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	db "3chan/models"

	validation "github.com/go-ozzo/ozzo-validation"
)

// GetThreadsList GetThreadsList
func GetThreadsList(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{"threads": db.FindALLThreads()})
}

// MakeThread MakeThread
func MakeThread(c *gin.Context) {
	threadName := c.PostForm("title")
	err := validation.Validate(threadName,
		validation.Required,
		validation.Length(1, 25),
	)
	if err != nil {
		log.Fatal(err)
	}
	db.InsertThread(threadName)
	c.HTML(http.StatusOK, "index.tmpl", gin.H{"threads": db.FindALLThreads()})
}

// GetThread GetThread
func GetThread(c *gin.Context) {
	threadID := c.Param("threadId")
	err := validation.Validate(threadID,
		validation.Required,
		validation.Length(1, 25),
	)
	if err != nil {
		log.Fatal(err)
	}
	c.HTML(http.StatusOK, "board.tmpl", gin.H{"resList": db.FindThread(threadID)})
}

// PostMessage PostMessage
func PostMessage(c *gin.Context) {
	threadID := c.Param("threadId")
	postViewName := c.PostForm("viewName")
	postMessage := c.PostForm("message")
	err := validation.Validate(postViewName,
		validation.Required,
		validation.Length(1, 25),
	)
	if err != nil {
		log.Fatal(err)
	}
	db.InsertMessage(threadID, postViewName, postMessage)
	c.HTML(http.StatusOK, "board.tmpl", gin.H{"resList": db.FindThread(threadID)})
}
