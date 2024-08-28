package interact

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/interact"
)

// AlibabaInteractSensorGcanvas gcanvas
// alibaba.interact.sensor.gcanvas
//
// gcanvas 功能
func AlibabaInteractSensorGcanvas(ctx context.Context, clt *core.SDKClient, req *interact.AlibabaInteractSensorGcanvasAPIRequest, resp *interact.AlibabaInteractSensorGcanvasAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
