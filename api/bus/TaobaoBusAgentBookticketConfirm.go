package bus

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/bus"
)

// TaobaoBusAgentBookticketConfirm 汽车票代理商接口—确认出票是否成功
// taobao.bus.agent.bookticket.confirm
//
// 代理商通过该接口通知汽车票系统订单出票结果。
func TaobaoBusAgentBookticketConfirm(ctx context.Context, clt *core.SDKClient, req *bus.TaobaoBusAgentBookticketConfirmAPIRequest, resp *bus.TaobaoBusAgentBookticketConfirmAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
