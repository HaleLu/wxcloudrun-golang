package model

import (
	"gorm.io/gorm"
	"time"
)

// Product 商品表
type Product struct {
	Id         uint64         `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"` // 商品id
	Content    string         `gorm:"column:content;NOT NULL" json:"content"`         // 商品内容
	Image      string         `gorm:"column:image" json:"image"`                      // 商品图片
	Price      uint64         `gorm:"column:price;NOT NULL" json:"price"`             // 商品积分
	CreateTime *time.Time     `gorm:"column:create_time;default:CURRENT_TIMESTAMP;NOT NULL" json:"create_time"`
	UpdateTime *time.Time     `gorm:"column:update_time;default:CURRENT_TIMESTAMP;NOT NULL" json:"update_time"`
	DeleteTime gorm.DeletedAt `gorm:"column:delete_time"`
}

func (m *Product) TableName() string {
	return "product"
}
