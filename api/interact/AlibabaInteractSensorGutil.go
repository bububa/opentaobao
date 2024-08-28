package interact

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/interact"
)

// AlibabaInteractSensorGutil canvas工具包
// alibaba.interact.sensor.gutil
//
// canvas工具包
func AlibabaInteractSensorGutil(ctx context.Context, clt *core.SDKClient, req *interact.AlibabaInteractSensorGutilAPIRequest, resp *interact.AlibabaInteractSensorGutilAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
