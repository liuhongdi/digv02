package global

import "github.com/jinzhu/gorm"

var (
	DBLink *gorm.DB
)

func SetupDBLink() error {
	var err error
	DBLink, err = gorm.Open("mysql", "root:password@tcp(127.0.0.1:3306)/dig?charset=utf8&parseTime=True&loc=Local")
	if err == nil  {
		// 全局禁用表名复数
		DBLink.SingularTable(true) // 如果设置为true,`User`的默认表名为`user`,使用`TableName`设置的表名不受影响
		//打开sql日志
		DBLink.LogMode(true)
		return nil
	} else {
		return err
	}
}
