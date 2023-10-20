package interact

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/interact"
)

// AlibabaInteractSensorTitlebarhide 隐藏titleBar
// alibaba.interact.sensor.titlebarhide
//
// 隐藏titleBar
func AlibabaInteractSensorTitlebarhide(clt *core.SDKClient, req *interact.AlibabaInteractSensorTitlebarhideAPIRequest, session string) (*interact.AlibabaInteractSensorTitlebarhideAPIResponse, error) {
	var resp interact.AlibabaInteractSensorTitlebarhideAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
