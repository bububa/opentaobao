package interact

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaInteractSensorVibrateAPIResponse 客户端震动 API返回值
// alibaba.interact.sensor.vibrate
//
// 客户端震动
type AlibabaInteractSensorVibrateAPIResponse struct {
	model.CommonResponse
	AlibabaInteractSensorVibrateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaInteractSensorVibrateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaInteractSensorVibrateAPIResponseModel).Reset()
}

// AlibabaInteractSensorVibrateAPIResponseModel is 客户端震动 成功返回结果
type AlibabaInteractSensorVibrateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_interact_sensor_vibrate_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result=0表示成功
	Result string `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaInteractSensorVibrateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = ""
}

var poolAlibabaInteractSensorVibrateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaInteractSensorVibrateAPIResponse)
	},
}

// GetAlibabaInteractSensorVibrateAPIResponse 从 sync.Pool 获取 AlibabaInteractSensorVibrateAPIResponse
func GetAlibabaInteractSensorVibrateAPIResponse() *AlibabaInteractSensorVibrateAPIResponse {
	return poolAlibabaInteractSensorVibrateAPIResponse.Get().(*AlibabaInteractSensorVibrateAPIResponse)
}

// ReleaseAlibabaInteractSensorVibrateAPIResponse 将 AlibabaInteractSensorVibrateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaInteractSensorVibrateAPIResponse(v *AlibabaInteractSensorVibrateAPIResponse) {
	v.Reset()
	poolAlibabaInteractSensorVibrateAPIResponse.Put(v)
}
