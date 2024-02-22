package net_edu

import (
	"context"

	"github.com/iooikaak/frame/gorm"
	model "github.com/iooikaak/job/common/model/mysql/net_edu"
)

func (d *Dao) GetNetEdu(ctx context.Context, id int64) (result *model.NetEdu, err error) {
	result = &model.NetEdu{}
	err = d.Db.Context(ctx).Model(model.NetEdu{}).Where("`id` = ?", id).First(result).Error
	return
}

func (d *Dao) GetNetEduN(db *gorm.DB, id int64) (result *model.NetEdu, err error) {
	result = &model.NetEdu{}
	err = db.Model(model.NetEdu{}).Where("`id` = ?", id).Order("id desc, name desc", true).Find(&result).Error
	return
}

func (d *Dao) UpdateNetEdu(db *gorm.DB, m, result *model.NetEdu) (err error) {
	if db == nil {
		return
	}
	err = db.Debug().Model(m).Limit(1).Update(&result).Error
	return
}

func (d *Dao) InsertNetEdu(db *gorm.DB, result *model.NetEdu) (err error) {
	if db == nil {
		return
	}
	//err = db.Omit("body").Create(&result).Error
	err = db.Create(&result).Error
	return
}

func (d *Dao) GetNetEdus(ctx context.Context, name string) (result []*model.NetEdu, err error) {
	err = d.Db.Context(ctx).Where("`name` = ?", name).Find(&result).Error
	return
}
