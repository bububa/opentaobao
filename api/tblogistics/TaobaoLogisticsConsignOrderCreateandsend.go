package tblogistics

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tblogistics"
)

// TaobaoLogisticsConsignOrderCreateandsend 创建订单并发货
// taobao.logistics.consign.order.createandsend
//
// 创建物流订单，并发货。
func TaobaoLogisticsConsignOrderCreateandsend(clt *core.SDKClient, req *tblogistics.TaobaoLogisticsConsignOrderCreateandsendAPIRequest, session string) (*tblogistics.TaobaoLogisticsConsignOrderCreateandsendAPIResponse, error) {
	var resp tblogistics.TaobaoLogisticsConsignOrderCreateandsendAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
