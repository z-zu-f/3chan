package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	db "3chan/models"
)

// GetThreadsList GetThreadsList
func GetThreadsList(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{"threads": db.FindALLThreads()})
}

// MakeThread MakeThread
func MakeThread(c *gin.Context) {
	threadName := c.PostForm("title")
	db.InsertThread(threadName)
	c.HTML(http.StatusOK, "index.tmpl", gin.H{"threads": db.FindALLThreads()})
}

// GetThread GetThread
func GetThread(c *gin.Context) {
	threadID := c.Param("threadId")
	c.HTML(http.StatusOK, "board.tmpl", gin.H{"resList": db.FindThread(threadID)})
}

// PostMessage PostMessage
func PostMessage(c *gin.Context) {
	threadID := c.Param("threadId")
	postViewName := c.PostForm("viewName")
	postMessage := c.PostForm("message")
	db.InsertMessage(threadID, postViewName, postMessage)
	c.HTML(http.StatusOK, "board.tmpl", gin.H{"resList": db.FindThread(threadID)})
}
