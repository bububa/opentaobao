package iot

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/iot"
)

// TaobaoAilabAicloudTopDeviceControlLamp 台灯控制
// taobao.ailab.aicloud.top.device.control.lamp
//
// 台灯控制
func TaobaoAilabAicloudTopDeviceControlLamp(clt *core.SDKClient, req *iot.TaobaoAilabAicloudTopDeviceControlLampAPIRequest, resp *iot.TaobaoAilabAicloudTopDeviceControlLampAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
