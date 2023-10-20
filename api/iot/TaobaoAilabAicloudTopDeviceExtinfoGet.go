package iot

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/iot"
)

// TaobaoAilabAicloudTopDeviceExtinfoGet 获取设备扩展信息
// taobao.ailab.aicloud.top.device.extinfo.get
//
// 获取设备扩展信息
func TaobaoAilabAicloudTopDeviceExtinfoGet(clt *core.SDKClient, req *iot.TaobaoAilabAicloudTopDeviceExtinfoGetAPIRequest, resp *iot.TaobaoAilabAicloudTopDeviceExtinfoGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
