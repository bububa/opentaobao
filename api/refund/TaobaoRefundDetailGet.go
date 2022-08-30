package refund

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/refund"
)

// TaobaoRefundDetailGet 退款详情页渲染
// taobao.refund.detail.get
//
// 退款详情页渲染
func TaobaoRefundDetailGet(clt *core.SDKClient, req *refund.TaobaoRefundDetailGetAPIRequest, session string) (*refund.TaobaoRefundDetailGetAPIResponse, error) {
	var resp refund.TaobaoRefundDetailGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
