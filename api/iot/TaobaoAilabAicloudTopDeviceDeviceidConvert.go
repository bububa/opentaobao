package iot

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/iot"
)

// TaobaoAilabAicloudTopDeviceDeviceidConvert 开放设备id转换内部设备id
// taobao.ailab.aicloud.top.device.deviceid.convert
//
// 将开放设备id转换为内部设备id
func TaobaoAilabAicloudTopDeviceDeviceidConvert(clt *core.SDKClient, req *iot.TaobaoAilabAicloudTopDeviceDeviceidConvertAPIRequest, resp *iot.TaobaoAilabAicloudTopDeviceDeviceidConvertAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
