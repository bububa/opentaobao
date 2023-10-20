package perfect

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/perfect"
)

// AlibabaTcwmsOutboundPickReceive 拣货接单
// alibaba.tcwms.outbound.pick.receive
//
// 拣货接单
func AlibabaTcwmsOutboundPickReceive(clt *core.SDKClient, req *perfect.AlibabaTcwmsOutboundPickReceiveAPIRequest, resp *perfect.AlibabaTcwmsOutboundPickReceiveAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
