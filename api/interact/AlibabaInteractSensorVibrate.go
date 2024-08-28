package interact

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/interact"
)

// AlibabaInteractSensorVibrate 客户端震动
// alibaba.interact.sensor.vibrate
//
// 客户端震动
func AlibabaInteractSensorVibrate(ctx context.Context, clt *core.SDKClient, req *interact.AlibabaInteractSensorVibrateAPIRequest, resp *interact.AlibabaInteractSensorVibrateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
