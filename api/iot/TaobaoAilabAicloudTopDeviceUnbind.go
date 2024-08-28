package iot

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/iot"
)

// TaobaoAilabAicloudTopDeviceUnbind 解绑设备
// taobao.ailab.aicloud.top.device.unbind
//
// 解绑设备
func TaobaoAilabAicloudTopDeviceUnbind(ctx context.Context, clt *core.SDKClient, req *iot.TaobaoAilabAicloudTopDeviceUnbindAPIRequest, resp *iot.TaobaoAilabAicloudTopDeviceUnbindAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
