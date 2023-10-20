package interact

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/interact"
)

// AlibabaInteractSensorOpenwindow 客户端打开新页面
// alibaba.interact.sensor.openwindow
//
// 客户端打开新页面
func AlibabaInteractSensorOpenwindow(clt *core.SDKClient, req *interact.AlibabaInteractSensorOpenwindowAPIRequest, resp *interact.AlibabaInteractSensorOpenwindowAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
