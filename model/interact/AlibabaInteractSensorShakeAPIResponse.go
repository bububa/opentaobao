package interact

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaInteractSensorShakeAPIResponse 摇一摇 API返回值
// alibaba.interact.sensor.shake
//
// 摇一摇
type AlibabaInteractSensorShakeAPIResponse struct {
	model.CommonResponse
	AlibabaInteractSensorShakeAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaInteractSensorShakeAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaInteractSensorShakeAPIResponseModel).Reset()
}

// AlibabaInteractSensorShakeAPIResponseModel is 摇一摇 成功返回结果
type AlibabaInteractSensorShakeAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_interact_sensor_shake_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 0
	Result string `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaInteractSensorShakeAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = ""
}

var poolAlibabaInteractSensorShakeAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaInteractSensorShakeAPIResponse)
	},
}

// GetAlibabaInteractSensorShakeAPIResponse 从 sync.Pool 获取 AlibabaInteractSensorShakeAPIResponse
func GetAlibabaInteractSensorShakeAPIResponse() *AlibabaInteractSensorShakeAPIResponse {
	return poolAlibabaInteractSensorShakeAPIResponse.Get().(*AlibabaInteractSensorShakeAPIResponse)
}

// ReleaseAlibabaInteractSensorShakeAPIResponse 将 AlibabaInteractSensorShakeAPIResponse 保存到 sync.Pool
func ReleaseAlibabaInteractSensorShakeAPIResponse(v *AlibabaInteractSensorShakeAPIResponse) {
	v.Reset()
	poolAlibabaInteractSensorShakeAPIResponse.Put(v)
}
