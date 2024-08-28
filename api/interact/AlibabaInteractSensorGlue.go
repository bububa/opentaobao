package interact

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/interact"
)

// AlibabaInteractSensorGlue 视频播放
// alibaba.interact.sensor.glue
//
// 视频播放
func AlibabaInteractSensorGlue(ctx context.Context, clt *core.SDKClient, req *interact.AlibabaInteractSensorGlueAPIRequest, resp *interact.AlibabaInteractSensorGlueAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
