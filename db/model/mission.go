package model

import (
	"gorm.io/gorm"
	"time"
)

var _ DbModel[uint64] = Mission{}

type MissionType uint8

const (
	MissionTypeOnce MissionType = iota
	MissionTypeEveryday
	MissionTypeWorkdays
	MissionTypeWeekends
)

// Mission 任务表
//CREATE TABLE `mission` (
// `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '任务id',
// `mission_type` tinyint(4) NOT NULL COMMENT '任务类型 0-一次性任务 1-每日任务 2-工作日任务 3-周末任务',
// `title` varchar(100) NOT NULL COMMENT '任务标题',
// `image` varchar(100) DEFAULT NULL COMMENT '任务图片',
// `price` bigint(20) unsigned NOT NULL COMMENT '任务价值',
// `last_finish_time` timestamp NULL DEFAULT NULL COMMENT '上次完成时间',
// `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
// `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
// `delete_time` timestamp NULL DEFAULT NULL,
// PRIMARY KEY (`id`)
//) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='任务表'
type Mission struct {
	Id             uint64         `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`   // 任务id
	MissionType    MissionType    `gorm:"column:mission_type;NOT NULL" json:"mission_type"` // 任务类型
	Title          string         `gorm:"column:title;NOT NULL" json:"title"`               // 任务标题
	Image          string         `gorm:"column:image" json:"image"`                        // 任务图片
	Price          uint64         `gorm:"column:price;NOT NULL" json:"price"`               // 任务价值
	LastFinishTime *time.Time     `gorm:"column:last_finish_time" json:"last_finish_time"`  // 上次完成时间
	CreateTime     *time.Time     `gorm:"column:create_time;default:CURRENT_TIMESTAMP;NOT NULL" json:"create_time"`
	UpdateTime     *time.Time     `gorm:"column:update_time;default:CURRENT_TIMESTAMP;NOT NULL" json:"update_time"`
	DeleteTime     gorm.DeletedAt `gorm:"column:delete_time"`
}

func (p Mission) GetId() uint64 {
	return p.Id
}

func (p Mission) SetId(id uint64) {
	p.Id = id
}

func (p Mission) TableName() string {
	return "mission"
}

func (p Mission) SetUpdateTime(time time.Time) {
	p.UpdateTime = &time
}
