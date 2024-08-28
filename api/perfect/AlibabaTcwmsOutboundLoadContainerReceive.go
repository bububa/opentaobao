package perfect

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/perfect"
)

// AlibabaTcwmsOutboundLoadContainerReceive 装箱接单
// alibaba.tcwms.outbound.load.container.receive
//
// 装箱接单
func AlibabaTcwmsOutboundLoadContainerReceive(ctx context.Context, clt *core.SDKClient, req *perfect.AlibabaTcwmsOutboundLoadContainerReceiveAPIRequest, resp *perfect.AlibabaTcwmsOutboundLoadContainerReceiveAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
