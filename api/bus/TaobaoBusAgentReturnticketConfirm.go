package bus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/bus"
)

// Taobaobusagentreturnticketconfirm 商家回调退票
// taobao.bus.agent.returnticket.confirm
//
// 商家通过TOP接口调用来回传退票状态
func Taobaobusagentreturnticketconfirm(clt *core.SDKClient, req *bus.TaobaobusagentreturnticketconfirmAPIRequest, session string) (*bus.TaobaobusagentreturnticketconfirmAPIResponse, error) {
	var resp bus.TaobaobusagentreturnticketconfirmAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
