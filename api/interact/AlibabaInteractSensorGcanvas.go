package interact

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/interact"
)

// Alibabainteractsensorgcanvas gcanvas
// alibaba.interact.sensor.gcanvas
//
// gcanvas 功能
func Alibabainteractsensorgcanvas(clt *core.SDKClient, req *interact.AlibabainteractsensorgcanvasAPIRequest, session string) (*interact.AlibabainteractsensorgcanvasAPIResponse, error) {
	var resp interact.AlibabainteractsensorgcanvasAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
