package work

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/work"
	workReq "github.com/flipped-aurora/gin-vue-admin/server/model/work/request"
)

type OrdersService struct {
}

// CreateOrders 创建Orders记录
// Author [piexlmax](https://github.com/piexlmax)
func (ordersService *OrdersService) CreateOrders(orders *work.Orders) (err error) {
	err = global.GVA_DB.Create(orders).Error
	return err
}

// DeleteOrders 删除Orders记录
// Author [piexlmax](https://github.com/piexlmax)
func (ordersService *OrdersService) DeleteOrders(orders work.Orders) (err error) {
	err = global.GVA_DB.Delete(&orders).Error
	return err
}

// DeleteOrdersByIds 批量删除Orders记录
// Author [piexlmax](https://github.com/piexlmax)
func (ordersService *OrdersService) DeleteOrdersByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]work.Orders{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateOrders 更新Orders记录
// Author [piexlmax](https://github.com/piexlmax)
func (ordersService *OrdersService) UpdateOrders(orders work.Orders) (err error) {
	err = global.GVA_DB.Save(&orders).Error
	return err
}

// GetOrders 根据id获取Orders记录
// Author [piexlmax](https://github.com/piexlmax)
func (ordersService *OrdersService) GetOrders(id uint) (orders work.Orders, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&orders).Error
	return
}

// GetOrders 根据id获取Order_Goods记录
// Author [piexlmax](https://github.com/piexlmax)
func (ordersService *OrdersService) GetOrderGoods(oid string) (goods []work.OrderGoods, err error) {
	err = global.GVA_DB.Where("order_no = ?", oid).Find(&goods).Error
	return
}

// GetOrdersInfoList 分页获取Orders记录
// Author [piexlmax](https://github.com/piexlmax)
func (ordersService *OrdersService) GetOrdersInfoList(info workReq.OrdersSearch) (list []work.Orders, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&work.Orders{})
	var orderss []work.Orders
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&orderss).Error
	return orderss, total, err
}
