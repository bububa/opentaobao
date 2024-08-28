package interact

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/interact"
)

// AlibabaInteractSensorAuthorize 客户端授权页
// alibaba.interact.sensor.authorize
//
// 客户端授权页
func AlibabaInteractSensorAuthorize(ctx context.Context, clt *core.SDKClient, req *interact.AlibabaInteractSensorAuthorizeAPIRequest, resp *interact.AlibabaInteractSensorAuthorizeAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
