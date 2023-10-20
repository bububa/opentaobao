package logistic

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// TaobaoWmsOrderWarehouseRouteGet 获取订单仓库路由信息
// taobao.wms.order.warehouse.route.get
//
// 获取订单仓库路由信息
func TaobaoWmsOrderWarehouseRouteGet(clt *core.SDKClient, req *logistic.TaobaoWmsOrderWarehouseRouteGetAPIRequest, resp *logistic.TaobaoWmsOrderWarehouseRouteGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
