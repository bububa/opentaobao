package tblogistics

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tblogistics"
)

// TaobaoLogisticsConsignOrderCreateandsend 创建订单并发货
// taobao.logistics.consign.order.createandsend
//
// 创建物流订单，并发货。
func TaobaoLogisticsConsignOrderCreateandsend(clt *core.SDKClient, req *tblogistics.TaobaoLogisticsConsignOrderCreateandsendAPIRequest, resp *tblogistics.TaobaoLogisticsConsignOrderCreateandsendAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
