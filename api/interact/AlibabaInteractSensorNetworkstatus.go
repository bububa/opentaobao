package interact

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/interact"
)

// AlibabaInteractSensorNetworkstatus 网络状态
// alibaba.interact.sensor.networkstatus
//
// 客户端网络状态
func AlibabaInteractSensorNetworkstatus(clt *core.SDKClient, req *interact.AlibabaInteractSensorNetworkstatusAPIRequest, resp *interact.AlibabaInteractSensorNetworkstatusAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
