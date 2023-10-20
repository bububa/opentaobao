package perfect

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/perfect"
)

// AlibabaTcwmsOutboundPickReceive 拣货接单
// alibaba.tcwms.outbound.pick.receive
//
// 拣货接单
func AlibabaTcwmsOutboundPickReceive(clt *core.SDKClient, req *perfect.AlibabaTcwmsOutboundPickReceiveAPIRequest, session string) (*perfect.AlibabaTcwmsOutboundPickReceiveAPIResponse, error) {
	var resp perfect.AlibabaTcwmsOutboundPickReceiveAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
