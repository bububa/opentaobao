package interact

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/interact"
)

// AlibabaInteractSensorGmedia gmedia
// alibaba.interact.sensor.gmedia
//
// 媒体功能
func AlibabaInteractSensorGmedia(ctx context.Context, clt *core.SDKClient, req *interact.AlibabaInteractSensorGmediaAPIRequest, resp *interact.AlibabaInteractSensorGmediaAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
