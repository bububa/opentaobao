package interact

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/interact"
)

// AlibabaInteractSensorGmedia gmedia
// alibaba.interact.sensor.gmedia
//
// 媒体功能
func AlibabaInteractSensorGmedia(clt *core.SDKClient, req *interact.AlibabaInteractSensorGmediaAPIRequest, resp *interact.AlibabaInteractSensorGmediaAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
