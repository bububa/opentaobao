package iot

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/iot"
)

// TaobaoAilabAicloudTopDeviceExtinfoGet 获取设备扩展信息
// taobao.ailab.aicloud.top.device.extinfo.get
//
// 获取设备扩展信息
func TaobaoAilabAicloudTopDeviceExtinfoGet(ctx context.Context, clt *core.SDKClient, req *iot.TaobaoAilabAicloudTopDeviceExtinfoGetAPIRequest, resp *iot.TaobaoAilabAicloudTopDeviceExtinfoGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
