package interact

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/interact"
)

// AlibabaInteractSensorShare 分享
// alibaba.interact.sensor.share
//
// 客户端分享
func AlibabaInteractSensorShare(clt *core.SDKClient, req *interact.AlibabaInteractSensorShareAPIRequest, resp *interact.AlibabaInteractSensorShareAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
