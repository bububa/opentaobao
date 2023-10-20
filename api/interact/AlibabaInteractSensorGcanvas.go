package interact

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/interact"
)

// AlibabaInteractSensorGcanvas gcanvas
// alibaba.interact.sensor.gcanvas
//
// gcanvas 功能
func AlibabaInteractSensorGcanvas(clt *core.SDKClient, req *interact.AlibabaInteractSensorGcanvasAPIRequest, resp *interact.AlibabaInteractSensorGcanvasAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
