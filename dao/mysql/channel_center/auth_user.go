package channel_center

//
//import (
//	"context"
//
//	model "git.juqitech.com/go/drama/common/model/mysql/channel_center"
//	"github.com/iooikaak/frame/gorm"
//)
//
//func (d *Dao) GetAuthUser(ctx context.Context, id int64) (result *model.AuthUser, err error) {
//	result = &model.AuthUser{}
//	err = d.Db.Context(ctx).Model(model.AuthUser{}).Where("`id` = ?", id).First(result).Error
//	return
//}
//
//func (d *Dao) GetAuthUserN(db *gorm.DB, id int64) (result *model.AuthUser, err error) {
//	result = &model.AuthUser{}
//	err = db.Model(model.AuthUser{}).Where("`id` = ?", id).Order("id desc, name desc", true).Find(&result).Error
//	return
//}
//
//func (d *Dao) UpdateAuthUser(db *gorm.DB, m, result *model.AuthUser) (err error) {
//	if db == nil {
//		return
//	}
//	err = db.Debug().Model(m).Limit(1).Update(&result).Error
//	return
//}
//
//func (d *Dao) InsertAuthUser(db *gorm.DB, result *model.AuthUser) (err error) {
//	if db == nil {
//		return
//	}
//	err = db.Omit("pwd").Create(&result).Error
//	return
//}
//
//func (d *Dao) GetAuthUsers(ctx context.Context, name string) (result []*model.AuthUser, err error) {
//	err = d.Db.Context(ctx).Where("`name` = ?", name).Find(&result).Error
//	return
//}
