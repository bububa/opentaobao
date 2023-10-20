package iot

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/iot"
)

// TaobaoAilabAicloudTopDeviceControlChildlock 设备儿童锁
// taobao.ailab.aicloud.top.device.control.childlock
//
// 设备儿童锁
func TaobaoAilabAicloudTopDeviceControlChildlock(clt *core.SDKClient, req *iot.TaobaoAilabAicloudTopDeviceControlChildlockAPIRequest, resp *iot.TaobaoAilabAicloudTopDeviceControlChildlockAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
