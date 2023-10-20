package interact

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/interact"
)

// AlibabaInteractSensorTitlebarhide 隐藏titleBar
// alibaba.interact.sensor.titlebarhide
//
// 隐藏titleBar
func AlibabaInteractSensorTitlebarhide(clt *core.SDKClient, req *interact.AlibabaInteractSensorTitlebarhideAPIRequest, resp *interact.AlibabaInteractSensorTitlebarhideAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
