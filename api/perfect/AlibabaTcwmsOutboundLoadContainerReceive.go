package perfect

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/perfect"
)

// AlibabaTcwmsOutboundLoadContainerReceive 装箱接单
// alibaba.tcwms.outbound.load.container.receive
//
// 装箱接单
func AlibabaTcwmsOutboundLoadContainerReceive(clt *core.SDKClient, req *perfect.AlibabaTcwmsOutboundLoadContainerReceiveAPIRequest, resp *perfect.AlibabaTcwmsOutboundLoadContainerReceiveAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
