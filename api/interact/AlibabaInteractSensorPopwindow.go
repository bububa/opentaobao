package interact

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/interact"
)

// AlibabaInteractSensorPopwindow popwindow
// alibaba.interact.sensor.popwindow
//
// popwindow
func AlibabaInteractSensorPopwindow(ctx context.Context, clt *core.SDKClient, req *interact.AlibabaInteractSensorPopwindowAPIRequest, resp *interact.AlibabaInteractSensorPopwindowAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
