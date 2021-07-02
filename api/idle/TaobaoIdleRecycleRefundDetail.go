package idle

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

// TaobaoIdleRecycleRefundDetail 闲鱼回收退款详情V2
// taobao.idle.recycle.refund.detail
//
// 回收订单退款详情，主要包括退款状态，超时时间，和同意退款的卖家退货地址信息
func TaobaoIdleRecycleRefundDetail(clt *core.SDKClient, req *idle.TaobaoIdleRecycleRefundDetailAPIRequest, session string) (*idle.TaobaoIdleRecycleRefundDetailAPIResponse, error) {
	var resp idle.TaobaoIdleRecycleRefundDetailAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
