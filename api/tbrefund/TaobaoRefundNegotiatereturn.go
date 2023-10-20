package tbrefund

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbrefund"
)

// Taobaorefundnegotiatereturn 协商退货退款
// taobao.refund.negotiatereturn
//
// 协商退货退款
func Taobaorefundnegotiatereturn(clt *core.SDKClient, req *tbrefund.TaobaorefundnegotiatereturnAPIRequest, session string) (*tbrefund.TaobaorefundnegotiatereturnAPIResponse, error) {
	var resp tbrefund.TaobaorefundnegotiatereturnAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
