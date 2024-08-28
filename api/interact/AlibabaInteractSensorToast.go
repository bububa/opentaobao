package interact

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/interact"
)

// AlibabaInteractSensorToast toast
// alibaba.interact.sensor.toast
//
// toast提示
func AlibabaInteractSensorToast(ctx context.Context, clt *core.SDKClient, req *interact.AlibabaInteractSensorToastAPIRequest, resp *interact.AlibabaInteractSensorToastAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
