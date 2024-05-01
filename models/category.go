package models

import (
	"github.com/provider-go/category/global"
	"time"
)

type Category struct {
	Id           int32     `json:"id" gorm:"auto_increment;primary_key;comment:'主键'"`
	ParentId     int32     `json:"parentId" gorm:"column:parent_id;not null;default:0;comment:父节点"`
	CategoryName string    `json:"categoryName" gorm:"column:category_name;type:varchar(20);not null;default:'';comment:类目名称"`
	Icon         string    `json:"icon" gorm:"column:icon;type:varchar(255);not null;default:'';comment:类目图标"`
	Pic          string    `json:"pic" gorm:"column:pic;type:varchar(255);not null;default:'';comment:类目的显示图片"`
	Brief        string    `json:"brief" gorm:"column:brief;type:varchar(255);not null;default:'';comment:简要描述"`
	Seq          int       `json:"seq" gorm:"column:seq;type:tinyint(1);not null;default:0;comment:排序"`
	Status       int       `json:"status" gorm:"column:status;type:tinyint(1);not null;default:0;comment:状态：0：正常；1：隐藏"`
	CreateTime   time.Time `json:"create_time" gorm:"autoCreateTime;comment:创建时间"`
	UpdateTime   time.Time `json:"update_time" gorm:"autoCreateTime;comment:更新时间"`
}

func Createcategory(parentId int32, categoryName, icon, pic, brief string) error {
	return global.DB.Table("categories").Select("parent_id", "category_name", "icon", "pic", "brief").
		Create(&Category{ParentId: parentId, CategoryName: categoryName, Icon: icon, Pic: pic, Brief: brief}).Error
}

func Deletecategory(id int32) error {
	return global.DB.Table("categories").Where("id = ?", id).Delete(&Category{}).Error
}

func Listcategory(upId, pageSize, pageNum int) ([]*Category, int64, error) {
	var rows []*Category
	//计算列表数量
	var count int64
	global.DB.Table("categories").Count(&count)

	if err := global.DB.Table("categories").Where("upid = ?", upId).Order("seq desc").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&rows).Error; err != nil {
		return nil, 0, err
	}
	return rows, count, nil
}

func Viewcategory(id int32) (*Category, error) {
	row := new(Category)
	if err := global.DB.Table("categories").Where("id = ?", id).First(&row).Error; err != nil {
		return nil, err
	}
	return row, nil
}
