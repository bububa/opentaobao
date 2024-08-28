package bus

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/bus"
)

// TaobaoBusAgentReturnticketConfirm 商家回调退票
// taobao.bus.agent.returnticket.confirm
//
// 商家通过TOP接口调用来回传退票状态
func TaobaoBusAgentReturnticketConfirm(ctx context.Context, clt *core.SDKClient, req *bus.TaobaoBusAgentReturnticketConfirmAPIRequest, resp *bus.TaobaoBusAgentReturnticketConfirmAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
