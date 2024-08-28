package iot

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/iot"
)

// TaobaoAilabAicloudTopDeviceAuthcodeGet 获取设备授权码
// taobao.ailab.aicloud.top.device.authcode.get
//
// 获取设备授权码
func TaobaoAilabAicloudTopDeviceAuthcodeGet(ctx context.Context, clt *core.SDKClient, req *iot.TaobaoAilabAicloudTopDeviceAuthcodeGetAPIRequest, resp *iot.TaobaoAilabAicloudTopDeviceAuthcodeGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
