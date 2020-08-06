package model

type User struct {
	Id        int64  `gorm:"primary_key;AUTO_INCREMENT;index;"` // 主键ID
	IdName    string `gorm:"size:255;"`                         // 身份证姓名
	IdCard    string `gorm:"size:255;index"`                    // 身份证号
	Mobile    string `gorm:"size:255;index;not null;"`          // 手机号
	Password  string `gorm:"size:255;"`                         // 密码
	HeadImg   string `gorm:"size:255;"`                         // 头像
	Level     int    `gorm:"DEFAULT:0;"`                        // 用户级别
	Status    int    `gorm:"DEFAULT:0;"`                        // 用户状态
	SerialNum string `gorm:"size:255;"`                         // 序列号

	NewcomerAward bool `gorm:"DEFAULT:0;"` // 是否领取新人奖

	CreateTime string // 创建时间
	ModifyTime string // 修改时间
}
