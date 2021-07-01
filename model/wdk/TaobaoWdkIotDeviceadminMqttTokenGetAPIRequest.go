package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoWdkIotDeviceadminMqttTokenGetAPIRequest
获取MQTT访问令牌 API请求
taobao.wdk.iot.deviceadmin.mqtt.token.get

智能硬件设备动态注册和获取mqtt设备信息 */
type TaobaoWdkIotDeviceadminMqttTokenGetAPIRequest struct {
	model.Params
	// accessKey
	_accessKey string
	// 申请令牌的客户端时间戳
	_applyTimestamp int64
}

// New
