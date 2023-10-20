package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWdkIotDeviceadminMqttTokenGetAPIResponse 获取MQTT访问令牌 API返回值
// taobao.wdk.iot.deviceadmin.mqtt.token.get
//
// 智能硬件设备动态注册和获取mqtt设备信息
type TaobaoWdkIotDeviceadminMqttTokenGetAPIResponse struct {
	model.CommonResponse
	TaobaoWdkIotDeviceadminMqttTokenGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoWdkIotDeviceadminMqttTokenGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoWdkIotDeviceadminMqttTokenGetAPIResponseModel).Reset()
}

// TaobaoWdkIotDeviceadminMqttTokenGetAPIResponseModel is 获取MQTT访问令牌 成功返回结果
type TaobaoWdkIotDeviceadminMqttTokenGetAPIResponseModel struct {
	XMLName xml.Name `xml:"wdk_iot_deviceadmin_mqtt_token_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TaobaoWdkIotDeviceadminMqttTokenGetHmResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoWdkIotDeviceadminMqttTokenGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoWdkIotDeviceadminMqttTokenGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoWdkIotDeviceadminMqttTokenGetAPIResponse)
	},
}

// GetTaobaoWdkIotDeviceadminMqttTokenGetAPIResponse 从 sync.Pool 获取 TaobaoWdkIotDeviceadminMqttTokenGetAPIResponse
func GetTaobaoWdkIotDeviceadminMqttTokenGetAPIResponse() *TaobaoWdkIotDeviceadminMqttTokenGetAPIResponse {
	return poolTaobaoWdkIotDeviceadminMqttTokenGetAPIResponse.Get().(*TaobaoWdkIotDeviceadminMqttTokenGetAPIResponse)
}

// ReleaseTaobaoWdkIotDeviceadminMqttTokenGetAPIResponse 将 TaobaoWdkIotDeviceadminMqttTokenGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoWdkIotDeviceadminMqttTokenGetAPIResponse(v *TaobaoWdkIotDeviceadminMqttTokenGetAPIResponse) {
	v.Reset()
	poolTaobaoWdkIotDeviceadminMqttTokenGetAPIResponse.Put(v)
}
