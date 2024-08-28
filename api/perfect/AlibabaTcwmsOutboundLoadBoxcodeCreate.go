package perfect

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/perfect"
)

// AlibabaTcwmsOutboundLoadBoxcodeCreate 创建箱号
// alibaba.tcwms.outbound.load.boxcode.create
//
// 创建箱号
func AlibabaTcwmsOutboundLoadBoxcodeCreate(ctx context.Context, clt *core.SDKClient, req *perfect.AlibabaTcwmsOutboundLoadBoxcodeCreateAPIRequest, resp *perfect.AlibabaTcwmsOutboundLoadBoxcodeCreateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
