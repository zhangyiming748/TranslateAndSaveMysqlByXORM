package model

import (
	"TranslateAndSaveMysqlByXORM/lib/mysql"
	"golang.org/x/exp/slog"
	"time"
)

type Word struct {
	Id         int       `xorm:"not null pk autoincr comment('主键id') INT(11)" json:"id"`
	ZhCn       string    `xorm:"comment('中文词') TEXT" json:"zh_cn"`
	Other      string    `xorm:"comment('其他语言') TEXT" json:"other"`
	Kind       string    `xorm:"comment('外文语种') TEXT" json:"kind"`
	UpdateTime time.Time `xorm:"updated comment('更新时间) DateTime" json:"update_time"`
	CreateTime time.Time `xorm:"created comment('创建时间') DateTime" json:"create_time"`
	DeleteTime time.Time `xorm:"deleted comment('创建时间') DateTime" json:"delete_time"`
}

func (w Word) FindByWord() (Word, bool, error) {
	model := Word{}
	has, err := mysql.GetEngine().NewSession().Where("other = ?", w.Other).Get(&model)
	if err != nil {
		return Word{}, has, err
	} else if !has {
		return Word{}, has, err
	} else {
		return model, has, err
	}
}

func (w Word) CreateOne() {
	insert, err := mysql.GetEngine().NewSession().Insert(&w)
	if err != nil {
		slog.Error("word表插入新数据失败", slog.Any("错误信息", err))
		return
	} else {
		slog.Debug("插入新数据成功", slog.Int64("条数", insert))
	}
}

func (w Word) DeleteAll() {
	i, err := mysql.GetEngine().NewSession().Where("1 = 1").Table("word").Delete(&Word{})
	if err != nil {
		return
	}
	slog.Info("删除记录", slog.Int64("成功条数", i))
}
