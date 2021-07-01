package logistic

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

/* TaobaoWmsOrderWarehouseRouteGet
获取订单仓库路由信息
taobao.wms.order.warehouse.route.get

获取订单仓库路由信息 */
func TaobaoWmsOrderWarehouseRouteGet(clt *core.SDKClient, req *logistic.TaobaoWmsOrderWarehouseRouteGetAPIRequest, session string) (*logistic.TaobaoWmsOrderWarehouseRouteGetAPIResponse, error) {
	var resp logistic.TaobaoWmsOrderWarehouseRouteGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
