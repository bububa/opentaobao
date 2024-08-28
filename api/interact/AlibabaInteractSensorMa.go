package interact

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/interact"
)

// AlibabaInteractSensorMa 码相关API
// alibaba.interact.sensor.ma
//
// 码相关API
func AlibabaInteractSensorMa(ctx context.Context, clt *core.SDKClient, req *interact.AlibabaInteractSensorMaAPIRequest, resp *interact.AlibabaInteractSensorMaAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
