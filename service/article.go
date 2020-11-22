package service

import (
	"fmt"
	"github.com/liuhongdi/digv02/dao"
	"github.com/liuhongdi/digv02/model"
)
//得到一篇文章的详情
func GetOneArticle(articleId int64) (*model.Article, error) {
	return dao.SelectOneArticle(articleId)
}

//得到多篇文章，按分页返回
func GetArticleList(page int ,pageSize int) ([]*model.Article,error) {
	articles, err := dao.SelectAllArticle(page,pageSize)
	if err != nil {
		fmt.Println("is error")
		return nil,err
	} else {
		fmt.Println("not is error")
		return articles,nil
	}
}