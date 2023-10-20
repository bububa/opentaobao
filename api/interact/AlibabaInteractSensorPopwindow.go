package interact

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/interact"
)

// AlibabaInteractSensorPopwindow popwindow
// alibaba.interact.sensor.popwindow
//
// popwindow
func AlibabaInteractSensorPopwindow(clt *core.SDKClient, req *interact.AlibabaInteractSensorPopwindowAPIRequest, resp *interact.AlibabaInteractSensorPopwindowAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
