package perfect

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/perfect"
)

// AlibabaTcwmsOutboundPickReceive 拣货接单
// alibaba.tcwms.outbound.pick.receive
//
// 拣货接单
func AlibabaTcwmsOutboundPickReceive(ctx context.Context, clt *core.SDKClient, req *perfect.AlibabaTcwmsOutboundPickReceiveAPIRequest, resp *perfect.AlibabaTcwmsOutboundPickReceiveAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
