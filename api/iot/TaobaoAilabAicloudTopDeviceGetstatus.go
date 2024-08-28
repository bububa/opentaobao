package iot

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/iot"
)

// TaobaoAilabAicloudTopDeviceGetstatus 获取设备状态
// taobao.ailab.aicloud.top.device.getstatus
//
// 获取设备状态
func TaobaoAilabAicloudTopDeviceGetstatus(ctx context.Context, clt *core.SDKClient, req *iot.TaobaoAilabAicloudTopDeviceGetstatusAPIRequest, resp *iot.TaobaoAilabAicloudTopDeviceGetstatusAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
