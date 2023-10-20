package interact

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/interact"
)

// AlibabaInteractSensorGravity 重力感应
// alibaba.interact.sensor.gravity
//
// native获取重力感应
func AlibabaInteractSensorGravity(clt *core.SDKClient, req *interact.AlibabaInteractSensorGravityAPIRequest, resp *interact.AlibabaInteractSensorGravityAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
