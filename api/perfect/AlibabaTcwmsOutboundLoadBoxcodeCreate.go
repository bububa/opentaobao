package perfect

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/perfect"
)

// AlibabaTcwmsOutboundLoadBoxcodeCreate 创建箱号
// alibaba.tcwms.outbound.load.boxcode.create
//
// 创建箱号
func AlibabaTcwmsOutboundLoadBoxcodeCreate(clt *core.SDKClient, req *perfect.AlibabaTcwmsOutboundLoadBoxcodeCreateAPIRequest, session string) (*perfect.AlibabaTcwmsOutboundLoadBoxcodeCreateAPIResponse, error) {
	var resp perfect.AlibabaTcwmsOutboundLoadBoxcodeCreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
