package model

import (
	"TranslateAndSaveMysqlByXORM/lib/mysql"
	"golang.org/x/exp/slog"
	"time"
)

type Word struct {
	Id         int       `xorm:"not null pk autoincr comment('主键id') INT(11)" json:"id"`
	ZhCn       string    `xorm:"comment('中文词') VARCHAR(255)" json:"zh_cn"`
	EnUs       string    `xorm:"comment('英文词') VARCHAR(255)" json:"en_us"`
	UpdateTime time.Time `xorm:"updated comment('更新时间) DateTime" json:"update_time"`
	CreateTime time.Time `xorm:"created comment('创建时间') DateTime" json:"create_time"`
	DeleteTime time.Time `xorm:"deleted comment('创建时间') DateTime" json:"delete_time"`
}

func (w Word) FindByEnglish() (Word, bool, error) {
	model := Word{}
	has, err := mysql.GetMysqlEngine().NewSession().Where("en_us =?", w.EnUs).Get(&model)
	if err != nil {
		return Word{}, has, err
	} else if !has {
		return Word{}, has, err
	} else {
		return model, has, err
	}
}
func (w Word) CreateOne() {
	//lib.InitMysql()
	mysql.GetMysqlEngine()
	insert, err := mysql.GetMysqlEngine().NewSession().Insert(&w)
	if err != nil {
		slog.Error("word表插入新数据失败", slog.Any("错误信息", err))
		return
	} else {
		slog.Debug("插入新数据成功", slog.Int64("条数", insert))
	}
}
