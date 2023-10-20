package iot

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/iot"
)

// TaobaoAilabAicloudTopDeviceControlVolume 设备音量
// taobao.ailab.aicloud.top.device.control.volume
//
// 设备音量
func TaobaoAilabAicloudTopDeviceControlVolume(clt *core.SDKClient, req *iot.TaobaoAilabAicloudTopDeviceControlVolumeAPIRequest, resp *iot.TaobaoAilabAicloudTopDeviceControlVolumeAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
