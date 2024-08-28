package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// TaobaoWdkIotDeviceadminMqttTokenGet 获取MQTT访问令牌
// taobao.wdk.iot.deviceadmin.mqtt.token.get
//
// 智能硬件设备动态注册和获取mqtt设备信息
func TaobaoWdkIotDeviceadminMqttTokenGet(ctx context.Context, clt *core.SDKClient, req *wdk.TaobaoWdkIotDeviceadminMqttTokenGetAPIRequest, resp *wdk.TaobaoWdkIotDeviceadminMqttTokenGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
