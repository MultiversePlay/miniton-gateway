package user

import (
	"context"
	"time"

	"www.miniton-gateway.com/pkg/mysql"
)

const TableName = "user"

type (
	Entity struct {
		ID         int64     `gorm:"column:id" json:"id"`
		TgID       int64     `gorm:"column:tg_id" json:"tg_id"`
		Name       string    `gorm:"column:name"  json:"name"`
		Region     int64     `gorm:"column:region" json:"region"`
		Gender     int       `gorm:"column:gender" json:"gender"`
		Birth      int       `gorm:"column:birth" json:"birth"`
		Avatar     string    `gorm:"column:avatar" json:"avatar"`
		Backgroud  string    `gorm:"column:backgroud" json:"backgroud"`
		CreateTime time.Time `gorm:"column:create_time" json:"create_time"`
		UpdateTime time.Time `gorm:"column:update_time" json:"update_time"`
	}
)

func (e *Entity) TableName() string {
	return TableName
}

func DetailByTGID(ctx context.Context, tgID int64) (entity *Entity, err error) {
	entity = new(Entity)
	err = mysql.DB.Model(entity).Where("tg_id = ?", tgID).Find(&entity).Error
	return
}

func Create(ctx context.Context, entity *Entity) (err error) {
	return mysql.DB.Model(&Entity{}).Create(entity).Error
}
