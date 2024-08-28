package interact

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/interact"
)

// AlibabaInteractSensorLogin 获取登陆页面
// alibaba.interact.sensor.login
//
// 获取登陆页面
func AlibabaInteractSensorLogin(ctx context.Context, clt *core.SDKClient, req *interact.AlibabaInteractSensorLoginAPIRequest, resp *interact.AlibabaInteractSensorLoginAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
