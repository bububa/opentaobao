package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoWdkIotDeviceadminMqttDeviceGetwithtokenAPIResponse
获取mqtt设备信息 API返回值
taobao.wdk.iot.deviceadmin.mqtt.device.getwithtoken

智能硬件设备动态注册和获取mqtt设备信息 */
type TaobaoWdkIotDeviceadminMqttDeviceGetwithtokenAPIResponse struct {
	model.CommonResponse
	TaobaoWdkIotDeviceadminMqttDeviceGetwithtokenAPIResponseModel
}

// TaobaoWdkIotDeviceadminMqttDeviceGetwithtokenAPIResponseModel is 获取mqtt设备信息 成功返回结果
type TaobaoWdkIotDeviceadminMqttDeviceGetwithtokenAPIResponseModel struct {
	XMLName xml.Name `xml:"wdk_iot_deviceadmin_mqtt_device_getwithtoken_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TaobaoWdkIotDeviceadminMqttDeviceGetwithtokenHmResult `json:"result,omitempty" xml:"result,omitempty"`
}
