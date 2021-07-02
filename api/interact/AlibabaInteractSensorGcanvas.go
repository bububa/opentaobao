package interact

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/interact"
)

// AlibabaInteractSensorGcanvas gcanvas
// alibaba.interact.sensor.gcanvas
//
// gcanvas 功能
func AlibabaInteractSensorGcanvas(clt *core.SDKClient, req *interact.AlibabaInteractSensorGcanvasAPIRequest, session string) (*interact.AlibabaInteractSensorGcanvasAPIResponse, error) {
	var resp interact.AlibabaInteractSensorGcanvasAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
