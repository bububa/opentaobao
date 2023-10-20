package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// TaobaoWdkIotDeviceadminMqttTokenGet 获取MQTT访问令牌
// taobao.wdk.iot.deviceadmin.mqtt.token.get
//
// 智能硬件设备动态注册和获取mqtt设备信息
func TaobaoWdkIotDeviceadminMqttTokenGet(clt *core.SDKClient, req *wdk.TaobaoWdkIotDeviceadminMqttTokenGetAPIRequest, session string) (*wdk.TaobaoWdkIotDeviceadminMqttTokenGetAPIResponse, error) {
	var resp wdk.TaobaoWdkIotDeviceadminMqttTokenGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
