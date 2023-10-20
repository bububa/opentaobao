package tbrefund

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbrefund"
)

// Taobaorefundnegotiatereturnrender 协商退货退款渲染
// taobao.refund.negotiatereturn.render
//
// 协商退货退款渲染
func Taobaorefundnegotiatereturnrender(clt *core.SDKClient, req *tbrefund.TaobaorefundnegotiatereturnrenderAPIRequest, session string) (*tbrefund.TaobaorefundnegotiatereturnrenderAPIResponse, error) {
	var resp tbrefund.TaobaorefundnegotiatereturnrenderAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
