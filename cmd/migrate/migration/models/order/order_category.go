// @Author sunwenbo
// 2024/7/16 21:14
package models

import (
	models2 "smart-api/cmd/migrate/migration/models"
)

type OrderCategory struct {
	ID            int             `gorm:"primaryKey;autoIncrement" json:"id"`
	Name          string          `gorm:"column:name;type:varchar(50)" json:"name"`
	Creator       string          `gorm:"creator:des;type:varchar(20)" json:"creator"`         // 创建人
	Regenerator   string          `gorm:"regenerator:des;type:varchar(20)" json:"regenerator"` // 更新人
	ChineseName   string          `gorm:"column:chineseName;type:varchar(50)" json:"chineseName"`
	FlowTemplates []FlowTemplates `gorm:"foreignKey:CategoryID" json:"flowTemplates"`
	models2.ControlBy
	models2.ModelTime
}

func (OrderCategory) TableName() string {
	return "order_category"
}
