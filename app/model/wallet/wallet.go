package wallet

import (
	"context"

	"www.miniton-gateway.com/pkg/mysql"
)

const TableName = "wallet"

type (
	Entity struct {
		ID          int64  `gorm:"column:id" json:"id"`
		TgID        int64  `gorm:"column:tg_id" json:"tg_id"`
		UserID      int64  `gorm:"column:user_id" json:"user_id"`
		TonAddress  string `gorm:"column:ton_address" json:"ton_address"`
		TonSeed     string `gorm:"column:ton_seed" json:"ton_seed"`
		TonPassword int64  `gorm:"column:ton_password" json:"ton_password"`
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
