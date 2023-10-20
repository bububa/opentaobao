package idle

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

// AlibabaIdleRentOrderQuery 查询订单
// alibaba.idle.rent.order.query
//
// 查询订单信息
func AlibabaIdleRentOrderQuery(clt *core.SDKClient, req *idle.AlibabaIdleRentOrderQueryAPIRequest, resp *idle.AlibabaIdleRentOrderQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
