package iot

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/iot"
)

// TaobaoAilabAicloudTopDeviceStatusinfoGet 获取设备状态信息
// taobao.ailab.aicloud.top.device.statusinfo.get
//
// 获取设备状态信息
func TaobaoAilabAicloudTopDeviceStatusinfoGet(ctx context.Context, clt *core.SDKClient, req *iot.TaobaoAilabAicloudTopDeviceStatusinfoGetAPIRequest, resp *iot.TaobaoAilabAicloudTopDeviceStatusinfoGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
