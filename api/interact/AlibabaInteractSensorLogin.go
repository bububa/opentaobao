package interact

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/interact"
)

// AlibabaInteractSensorLogin 获取登陆页面
// alibaba.interact.sensor.login
//
// 获取登陆页面
func AlibabaInteractSensorLogin(clt *core.SDKClient, req *interact.AlibabaInteractSensorLoginAPIRequest, resp *interact.AlibabaInteractSensorLoginAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
