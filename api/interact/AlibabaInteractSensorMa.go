package interact

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/interact"
)

// AlibabaInteractSensorMa 码相关API
// alibaba.interact.sensor.ma
//
// 码相关API
func AlibabaInteractSensorMa(clt *core.SDKClient, req *interact.AlibabaInteractSensorMaAPIRequest, resp *interact.AlibabaInteractSensorMaAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
