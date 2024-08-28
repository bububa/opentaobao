package interact

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/interact"
)

// AlibabaInteractSensorBlow 吹气
// alibaba.interact.sensor.blow
//
// 客户端吹气
func AlibabaInteractSensorBlow(ctx context.Context, clt *core.SDKClient, req *interact.AlibabaInteractSensorBlowAPIRequest, resp *interact.AlibabaInteractSensorBlowAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
