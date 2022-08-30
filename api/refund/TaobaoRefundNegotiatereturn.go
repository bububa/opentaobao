package refund

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/refund"
)

// TaobaoRefundNegotiatereturn 协商退货退款
// taobao.refund.negotiatereturn
//
// 协商退货退款
func TaobaoRefundNegotiatereturn(clt *core.SDKClient, req *refund.TaobaoRefundNegotiatereturnAPIRequest, session string) (*refund.TaobaoRefundNegotiatereturnAPIResponse, error) {
	var resp refund.TaobaoRefundNegotiatereturnAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
