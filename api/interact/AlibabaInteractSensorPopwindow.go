package interact

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/interact"
)

// AlibabaInteractSensorPopwindow popwindow
// alibaba.interact.sensor.popwindow
//
// popwindow
func AlibabaInteractSensorPopwindow(clt *core.SDKClient, req *interact.AlibabaInteractSensorPopwindowAPIRequest, session string) (*interact.AlibabaInteractSensorPopwindowAPIResponse, error) {
	var resp interact.AlibabaInteractSensorPopwindowAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
