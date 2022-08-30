package perfect

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/perfect"
)

// AlibabaTcwmsOutboundLoadContainerReceive 装箱接单
// alibaba.tcwms.outbound.load.container.receive
//
// 装箱接单
func AlibabaTcwmsOutboundLoadContainerReceive(clt *core.SDKClient, req *perfect.AlibabaTcwmsOutboundLoadContainerReceiveAPIRequest, session string) (*perfect.AlibabaTcwmsOutboundLoadContainerReceiveAPIResponse, error) {
	var resp perfect.AlibabaTcwmsOutboundLoadContainerReceiveAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
