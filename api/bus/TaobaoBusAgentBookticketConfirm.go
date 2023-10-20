package bus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/bus"
)

// Taobaobusagentbookticketconfirm 汽车票代理商接口—确认出票是否成功
// taobao.bus.agent.bookticket.confirm
//
// 代理商通过该接口通知汽车票系统订单出票结果。
func Taobaobusagentbookticketconfirm(clt *core.SDKClient, req *bus.TaobaobusagentbookticketconfirmAPIRequest, session string) (*bus.TaobaobusagentbookticketconfirmAPIResponse, error) {
	var resp bus.TaobaobusagentbookticketconfirmAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
