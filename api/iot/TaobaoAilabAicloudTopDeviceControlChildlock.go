package iot

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/iot"
)

// TaobaoAilabAicloudTopDeviceControlChildlock 设备儿童锁
// taobao.ailab.aicloud.top.device.control.childlock
//
// 设备儿童锁
func TaobaoAilabAicloudTopDeviceControlChildlock(ctx context.Context, clt *core.SDKClient, req *iot.TaobaoAilabAicloudTopDeviceControlChildlockAPIRequest, resp *iot.TaobaoAilabAicloudTopDeviceControlChildlockAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
