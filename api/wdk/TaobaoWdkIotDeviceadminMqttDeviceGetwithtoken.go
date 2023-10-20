package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// TaobaoWdkIotDeviceadminMqttDeviceGetwithtoken 获取mqtt设备信息
// taobao.wdk.iot.deviceadmin.mqtt.device.getwithtoken
//
// 智能硬件设备动态注册和获取mqtt设备信息
func TaobaoWdkIotDeviceadminMqttDeviceGetwithtoken(clt *core.SDKClient, req *wdk.TaobaoWdkIotDeviceadminMqttDeviceGetwithtokenAPIRequest, resp *wdk.TaobaoWdkIotDeviceadminMqttDeviceGetwithtokenAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
