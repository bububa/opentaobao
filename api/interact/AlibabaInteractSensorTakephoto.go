package interact

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/interact"
)

// AlibabaInteractSensorTakephoto takePhoto
// alibaba.interact.sensor.takephoto
//
// 客户端takePhoto
func AlibabaInteractSensorTakephoto(ctx context.Context, clt *core.SDKClient, req *interact.AlibabaInteractSensorTakephotoAPIRequest, resp *interact.AlibabaInteractSensorTakephotoAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
