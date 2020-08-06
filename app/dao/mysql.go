package dao

import (
	"ginpher/app/model"
	"github.com/vgmdj/utils/logger"
)

func (d *Dao) AutoMigrate() {
	d.db.AutoMigrate(&model.User{})
}

// Insert
func (d *Dao) Create(table string, value interface{}) error {
	err := d.db.Table(table).Create(value).Error

	if err != nil {
		logger.Error(err)
		return err
	}

	return nil
}

// Update
func (d *Dao) Update(table, where string, fields string, value interface{}) error {

	err := d.db.Table(table).Model(value).Where(where).Update(fields).Error
	if err != nil {
		logger.Error(err)
		return err
	}
	return nil
}

// Delete
// 删除
func (d *Dao) Delete(table string, where interface{}) error {

	err := d.db.Table(table).Delete(where).Error
	if err != nil {
		logger.Error(err)
		return err
	}

	return nil
}

// Find
// 获取所有匹配记录
func (d *Dao) Find(table, where, order string) (err error, count int64, mp interface{}) {

	err = d.db.Table(table).Where(where).Order(order).Find(mp).Count(count).Error
	if err != nil {
		logger.Error(err)
		return
	}

	return
}

// First
// 获取第一个匹配记录
func (d *Dao) First(table, where string) (err error, mp interface{}) {

	err = d.db.Table(table).Where(where).First(mp).Error
	if err != nil {
		logger.Error(err)
		return
	}

	return
}

// Last
//获取最后一条记录，按主键排序
func (d *Dao) Last(table string) (err error, mp interface{}) {

	err = d.db.Table(table).Last(mp).Error
	if err != nil {
		logger.Error(err)
		return
	}

	return
}

// Join
func (d *Dao) Join(tableMain, tableVice, fields, condition string) (err error, rows interface{}) {

	rows, err = d.db.Table(tableMain).Select(fields).Joins("left join" + tableVice + "on" + condition).Rows()
	if err != nil {
		logger.Error(err)
		return
	}

	return
}
