package interact

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/interact"
)

// AlibabaInteractSensorNetworkstatus 网络状态
// alibaba.interact.sensor.networkstatus
//
// 客户端网络状态
func AlibabaInteractSensorNetworkstatus(ctx context.Context, clt *core.SDKClient, req *interact.AlibabaInteractSensorNetworkstatusAPIRequest, resp *interact.AlibabaInteractSensorNetworkstatusAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
