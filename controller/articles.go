package controller

import (
	"net/http"
	"strconv"

	"goblog/dao"
	"goblog/server"

	"github.com/gin-gonic/gin"
)

func articleList(c *gin.Context, tab string, tpl string, page string, where map[string]interface{}) {
	// page := c.Query("page")
	if page == "" {
		page = "1"
	}
	pageNum, _ := strconv.ParseInt(page, 10, 32)
	pageNum = pageNum - 1
	list, _ := dao.GetArticlesList(int(pageNum), 15, where, "")
	baseinfo, _ := server.Getinfo()

	c.HTML(http.StatusOK, tpl, gin.H{
		"tab":         tab,
		"page":        pageNum + 2,
		"base":        baseinfo,
		"articleList": list,
	})
}

func Article(c *gin.Context) {
	page := c.Query("page")
	where := map[string]interface{}{
		"status = ?": 1,
	}
	articleList(c, "articles", "articlesPage.html", page, where)
	return
}

func Events(c *gin.Context) {
	page := c.Query("page")

	where := map[string]interface{}{
		"status = ?": 1,
	}
	articleList(c, "events", "events.html", page, where)
	return
}

func DaoJobs(c *gin.Context) {
	page := c.Query("page")

	where := map[string]interface{}{
		"status = ?": 1,
	}
	articleList(c, "jobs", "jobs.html", page, where)
	return
}
