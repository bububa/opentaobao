package omniorder

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/omniorder"
)

// TaobaoOmniOrderDetail 全渠道订单详情
// taobao.omni.order.detail
//
// 全渠道订单详情
func TaobaoOmniOrderDetail(clt *core.SDKClient, req *omniorder.TaobaoOmniOrderDetailAPIRequest, session string) (*omniorder.TaobaoOmniOrderDetailAPIResponse, error) {
	var resp omniorder.TaobaoOmniOrderDetailAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
