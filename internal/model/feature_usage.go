// Package model
// Wrote by yijian on 2024/08/16
package model

import (
	"time"
)

type FeatureUsage struct {
	ID        uint      `gorm:"column:f_id;primaryKey;autoIncrement"`
	UserIp    string    `gorm:"column:f_user_ip"`
	Feature   string    `gorm:"column:f_feature"`
	CreatedAt time.Time `gorm:"column:f_create_at"`
}
