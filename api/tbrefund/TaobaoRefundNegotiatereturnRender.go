package tbrefund

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbrefund"
)

// TaobaoRefundNegotiatereturnRender 协商退货退款渲染
// taobao.refund.negotiatereturn.render
//
// 协商退货退款渲染
func TaobaoRefundNegotiatereturnRender(clt *core.SDKClient, req *tbrefund.TaobaoRefundNegotiatereturnRenderAPIRequest, session string) (*tbrefund.TaobaoRefundNegotiatereturnRenderAPIResponse, error) {
	var resp tbrefund.TaobaoRefundNegotiatereturnRenderAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
