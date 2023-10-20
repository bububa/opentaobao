package bus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/bus"
)

// Taobaobusagentrefundticketconfirm 商家top回调退款明细
// taobao.bus.agent.refundticket.confirm
//
// 商家通过top回调告知平台退款明细
func Taobaobusagentrefundticketconfirm(clt *core.SDKClient, req *bus.TaobaobusagentrefundticketconfirmAPIRequest, session string) (*bus.TaobaobusagentrefundticketconfirmAPIResponse, error) {
	var resp bus.TaobaobusagentrefundticketconfirmAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
