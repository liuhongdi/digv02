package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/liuhongdi/digv02/dao"
	"github.com/liuhongdi/digv02/global"
	"github.com/liuhongdi/digv02/model"
	"github.com/liuhongdi/digv02/service"
	"strconv"
)

type ArticleController struct{}

func NewArticleController() ArticleController {
	return ArticleController{}
}
//得到一篇文章的详情
func (a *ArticleController) GetOne(c *gin.Context) {
    result := global.NewResult(c)
	id := c.Params.ByName("id")
	fmt.Println("id:"+id);

	articleId,err := strconv.ParseInt(id, 10, 64);
    if (err != nil) {
		result.Error(400,"参数错误")
		fmt.Println(err.Error())
		return
	}

	if (articleId == 100) {
		var z int = 0
		var i int = 100 / z
		fmt.Println("i:%i",i)
	}

	articleOne,err := service.GetOneArticle(articleId);
	if err != nil {
		result.Error(404,"数据查询错误")
	} else {
		result.Success(&articleOne);
	}
	return
}

//得到多篇文章，按分页返回
func (a *ArticleController) GetList(c *gin.Context) {
	result := global.NewResult(c)
	page := c.DefaultQuery("page", "1")
	pageInt, err := strconv.Atoi(page)
	if (err != nil) {
		//c.AbortWithStatus(400)
		result.Error(400,"参数错误")
		fmt.Println(err.Error())
		return
	}
	pageSize := 2;
	pageOffset := (pageInt-1) * pageSize

	articles,err := service.GetArticleList(pageOffset,pageSize)
	if err != nil {
		//c.AbortWithStatus(404)
		result.Error(404,"数据查询错误");
		fmt.Println(err.Error())
	} else {
		sum,_ := dao.SelectcountAll()
		pageInfo,_ := model.GetPageInfo(pageInt,pageSize,sum)
		result.Success(gin.H{"list":&articles,"pageinfo":pageInfo});
	}
	return
}

