package perfect

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/perfect"
)

// AlibabaTcwmsOutboundLoadBoxcodeCreate 创建箱号
// alibaba.tcwms.outbound.load.boxcode.create
//
// 创建箱号
func AlibabaTcwmsOutboundLoadBoxcodeCreate(clt *core.SDKClient, req *perfect.AlibabaTcwmsOutboundLoadBoxcodeCreateAPIRequest, resp *perfect.AlibabaTcwmsOutboundLoadBoxcodeCreateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
