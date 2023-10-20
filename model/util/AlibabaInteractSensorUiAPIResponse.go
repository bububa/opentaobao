package util

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaInteractSensorUiAPIResponse 基本ui操作 API返回值
// alibaba.interact.sensor.ui
//
// Weex 基本UI操作
type AlibabaInteractSensorUiAPIResponse struct {
	model.CommonResponse
	AlibabaInteractSensorUiAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaInteractSensorUiAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaInteractSensorUiAPIResponseModel).Reset()
}

// AlibabaInteractSensorUiAPIResponseModel is 基本ui操作 成功返回结果
type AlibabaInteractSensorUiAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_interact_sensor_ui_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 仅作客户端鉴权使用，不会发送接收请求
	Unnamed string `json:"unnamed,omitempty" xml:"unnamed,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaInteractSensorUiAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Unnamed = ""
}

var poolAlibabaInteractSensorUiAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaInteractSensorUiAPIResponse)
	},
}

// GetAlibabaInteractSensorUiAPIResponse 从 sync.Pool 获取 AlibabaInteractSensorUiAPIResponse
func GetAlibabaInteractSensorUiAPIResponse() *AlibabaInteractSensorUiAPIResponse {
	return poolAlibabaInteractSensorUiAPIResponse.Get().(*AlibabaInteractSensorUiAPIResponse)
}

// ReleaseAlibabaInteractSensorUiAPIResponse 将 AlibabaInteractSensorUiAPIResponse 保存到 sync.Pool
func ReleaseAlibabaInteractSensorUiAPIResponse(v *AlibabaInteractSensorUiAPIResponse) {
	v.Reset()
	poolAlibabaInteractSensorUiAPIResponse.Put(v)
}
