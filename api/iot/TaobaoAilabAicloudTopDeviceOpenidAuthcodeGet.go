package iot

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/iot"
)

// TaobaoAilabAicloudTopDeviceOpenidAuthcodeGet 获取openid设备通用授权码
// taobao.ailab.aicloud.top.device.openid.authcode.get
//
// 获取openid设备通用授权码
func TaobaoAilabAicloudTopDeviceOpenidAuthcodeGet(ctx context.Context, clt *core.SDKClient, req *iot.TaobaoAilabAicloudTopDeviceOpenidAuthcodeGetAPIRequest, resp *iot.TaobaoAilabAicloudTopDeviceOpenidAuthcodeGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
