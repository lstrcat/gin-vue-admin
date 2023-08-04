// 自动生成模板Orders
package work

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Orders 结构体
type Orders struct {
	global.GVA_MODEL
	Userid           *int       `json:"userid" form:"userid" gorm:"column:userid;comment:;"`
	Typecate         *int       `json:"typecate" form:"typecate" gorm:"column:typecate;comment:;"`
	Status           *int       `json:"status" form:"status" gorm:"column:status;comment:;"`
	Storename        string     `json:"storename" form:"storename" gorm:"column:storename;comment:;size:128;"`
	Paymode          string     `json:"paymode" form:"paymode" gorm:"column:paymode;comment:;size:255;"`
	Amount           string     `json:"amount" form:"amount" gorm:"column:amount;comment:;size:255;"`
	OrderNo          string     `json:"orderNo" form:"orderNo" gorm:"column:order_no;comment:;size:128;"`
	Postscript       string     `json:"postscript" form:"postscript" gorm:"column:postscript;comment:;size:255;"`
	ProductionedTime *time.Time `json:"productionedTime" form:"productionedTime" gorm:"column:productioned_time;comment:完成时间;"`
	AcceptName       string     `json:"acceptName" form:"acceptName" gorm:"column:accept_name;comment:;size:255;"`
	Mobile           string     `json:"mobile" form:"mobile" gorm:"column:mobile;comment:;size:255;"`
	Street           string     `json:"street" form:"street" gorm:"column:street;comment:;size:255;"`
	DoorNumber       string     `json:"doorNumber" form:"doorNumber" gorm:"column:door_number;comment:;size:128;"`
	CreatedTime      *time.Time `json:"createdTime" form:"createdTime" gorm:"column:created_time;comment:创建时间;"`
}

// TableName Orders 表名
func (Orders) TableName() string {
	return "orders"
}
