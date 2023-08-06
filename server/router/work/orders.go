package work

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type OrdersRouter struct {
}

// InitOrdersRouter 初始化 Orders 路由信息
func (s *OrdersRouter) InitOrdersRouter(Router *gin.RouterGroup) {
	ordersRouter := Router.Group("orders").Use(middleware.OperationRecord())
	ordersRouterWithoutRecord := Router.Group("orders")
	var ordersApi = v1.ApiGroupApp.WorkApiGroup.OrdersApi
	{
		ordersRouter.POST("createOrders", ordersApi.CreateOrders)             // 新建Orders
		ordersRouter.DELETE("deleteOrders", ordersApi.DeleteOrders)           // 删除Orders
		ordersRouter.DELETE("deleteOrdersByIds", ordersApi.DeleteOrdersByIds) // 批量删除Orders
		ordersRouter.PUT("updateOrders", ordersApi.UpdateOrders)              // 更新Orders
	}
	{
		ordersRouterWithoutRecord.GET("findOrders", ordersApi.FindOrders)         // 根据ID获取Orders
		ordersRouterWithoutRecord.GET("getOrdersList", ordersApi.GetOrdersList)   // 获取Orders列表
		ordersRouterWithoutRecord.GET("getOrdersCount", ordersApi.GetOrdersCount) // 获取Orders数量
	}
}
