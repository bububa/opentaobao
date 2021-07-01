package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoWmsOrderWarehouseRouteGetAPIRequest
获取订单仓库路由信息 API请求
taobao.wms.order.warehouse.route.get

获取订单仓库路由信息 */
type TaobaoWmsOrderWarehouseRouteGetAPIRequest struct {
	model.Params
	// 订单编号
	_orderCode string
}

// New
