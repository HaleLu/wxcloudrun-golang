package model

import (
	"gorm.io/gorm"
	"time"
)

var _ DbModel[uint64] = Product{}

// Product 商品表
// CREATE TABLE `product` (
//  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '商品id',
//  `content` varchar(100) NOT NULL COMMENT '商品内容',
//  `image` varchar(100) DEFAULT NULL COMMENT '商品图片',
//  `price` bigint(20) unsigned NOT NULL COMMENT '商品积分',
//  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
//  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
//  `delete_time` timestamp NULL DEFAULT NULL,
//  PRIMARY KEY (`id`)
//) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='商品表'
type Product struct {
	Id         uint64         `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"` // 商品id
	Content    string         `gorm:"column:content;NOT NULL" json:"content"`         // 商品内容
	Image      string         `gorm:"column:image" json:"image"`                      // 商品图片
	Price      uint64         `gorm:"column:price;NOT NULL" json:"price"`             // 商品积分
	CreateTime *time.Time     `gorm:"column:create_time;default:CURRENT_TIMESTAMP;NOT NULL" json:"create_time"`
	UpdateTime *time.Time     `gorm:"column:update_time;default:CURRENT_TIMESTAMP;NOT NULL" json:"update_time"`
	DeleteTime gorm.DeletedAt `gorm:"column:delete_time"`
}

func (p Product) GetId() uint64 {
	return p.Id
}

func (p Product) SetId(id uint64) {
	p.Id = id
}

func (p Product) TableName() string {
	return "product"
}

func (p Product) SetUpdateTime(time time.Time) {
	p.UpdateTime = &time
}
