package iot

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/iot"
)

// TaobaoAilabAicloudTopDeviceControlCustom 设备控制自定义扩展接口
// taobao.ailab.aicloud.top.device.control.custom
//
// 设备控制自定义扩展接口
func TaobaoAilabAicloudTopDeviceControlCustom(ctx context.Context, clt *core.SDKClient, req *iot.TaobaoAilabAicloudTopDeviceControlCustomAPIRequest, resp *iot.TaobaoAilabAicloudTopDeviceControlCustomAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
