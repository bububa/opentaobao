package tbrefund

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbrefund"
)

// TaobaoRefundNegotiatereturn 协商退货退款
// taobao.refund.negotiatereturn
//
// 协商退货退款
func TaobaoRefundNegotiatereturn(clt *core.SDKClient, req *tbrefund.TaobaoRefundNegotiatereturnAPIRequest, resp *tbrefund.TaobaoRefundNegotiatereturnAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
