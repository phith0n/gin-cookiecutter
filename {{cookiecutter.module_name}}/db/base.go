package db

import "time"

type BaseModel struct {
	ID             uint      `json:"id" gorm:"column:id;not null;primaryKey;autoIncrement;"`
	CreatedTime    time.Time `json:"created_time" gorm:"column:created_time;type:timestamptz(6);autoCreateTime;"`
	LastModifyTime time.Time `json:"last_modify_time" gorm:"column:last_modify_time;type:timestamptz(6);autoUpdateTime;"`
}
