package iot

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/iot"
)

// TaobaoAilabAicloudTopDeviceDetailinfoGet 获取设备详细信息
// taobao.ailab.aicloud.top.device.detailinfo.get
//
// 获取设备详细信息
func TaobaoAilabAicloudTopDeviceDetailinfoGet(ctx context.Context, clt *core.SDKClient, req *iot.TaobaoAilabAicloudTopDeviceDetailinfoGetAPIRequest, resp *iot.TaobaoAilabAicloudTopDeviceDetailinfoGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
