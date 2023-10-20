package logistic

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// Taobaowmsorderwarehouserouteget 获取订单仓库路由信息
// taobao.wms.order.warehouse.route.get
//
// 获取订单仓库路由信息
func Taobaowmsorderwarehouserouteget(clt *core.SDKClient, req *logistic.TaobaowmsorderwarehouseroutegetAPIRequest, session string) (*logistic.TaobaowmsorderwarehouseroutegetAPIResponse, error) {
	var resp logistic.TaobaowmsorderwarehouseroutegetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
