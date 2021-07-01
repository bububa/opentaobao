package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoWdkIotDeviceadminMqttDeviceGetwithtokenAPIRequest
获取mqtt设备信息 API请求
taobao.wdk.iot.deviceadmin.mqtt.device.getwithtoken

智能硬件设备动态注册和获取mqtt设备信息 */
type TaobaoWdkIotDeviceadminMqttDeviceGetwithtokenAPIRequest struct {
	model.Params
	// 业务编码，具体编码请联系杜尘
	_businessCode int64
	// 设备类型编码，具体编码请联系杜尘
	_deviceType int64
	// 环境编码，0日常，1预发，2线上
	_enviroCode int64
	// 设备的唯一标识码，比如网卡的MAC地址
	_deviceId string
	// 访问令牌
	_applyAccessToken string
}

// New
