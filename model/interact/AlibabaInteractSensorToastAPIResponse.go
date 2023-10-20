package interact

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaInteractSensorToastAPIResponse toast API返回值
// alibaba.interact.sensor.toast
//
// toast提示
type AlibabaInteractSensorToastAPIResponse struct {
	model.CommonResponse
	AlibabaInteractSensorToastAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaInteractSensorToastAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaInteractSensorToastAPIResponseModel).Reset()
}

// AlibabaInteractSensorToastAPIResponseModel is toast 成功返回结果
type AlibabaInteractSensorToastAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_interact_sensor_toast_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result=0 表示成功
	Result string `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaInteractSensorToastAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = ""
}

var poolAlibabaInteractSensorToastAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaInteractSensorToastAPIResponse)
	},
}

// GetAlibabaInteractSensorToastAPIResponse 从 sync.Pool 获取 AlibabaInteractSensorToastAPIResponse
func GetAlibabaInteractSensorToastAPIResponse() *AlibabaInteractSensorToastAPIResponse {
	return poolAlibabaInteractSensorToastAPIResponse.Get().(*AlibabaInteractSensorToastAPIResponse)
}

// ReleaseAlibabaInteractSensorToastAPIResponse 将 AlibabaInteractSensorToastAPIResponse 保存到 sync.Pool
func ReleaseAlibabaInteractSensorToastAPIResponse(v *AlibabaInteractSensorToastAPIResponse) {
	v.Reset()
	poolAlibabaInteractSensorToastAPIResponse.Put(v)
}
