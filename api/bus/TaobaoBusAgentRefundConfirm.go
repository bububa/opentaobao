package bus

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/bus"
)

// TaobaoBusAgentRefundConfirm 汽车票退票和退款二合一接口
// taobao.bus.agent.refund.confirm
//
// 1.商家退票成功后，回调飞猪平台汽车票退票接口，平台进行退票和退款操作。
func TaobaoBusAgentRefundConfirm(ctx context.Context, clt *core.SDKClient, req *bus.TaobaoBusAgentRefundConfirmAPIRequest, resp *bus.TaobaoBusAgentRefundConfirmAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
