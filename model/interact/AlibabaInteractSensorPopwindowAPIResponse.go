package interact

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaInteractSensorPopwindowAPIResponse popwindow API返回值
// alibaba.interact.sensor.popwindow
//
// popwindow
type AlibabaInteractSensorPopwindowAPIResponse struct {
	model.CommonResponse
	AlibabaInteractSensorPopwindowAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaInteractSensorPopwindowAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaInteractSensorPopwindowAPIResponseModel).Reset()
}

// AlibabaInteractSensorPopwindowAPIResponseModel is popwindow 成功返回结果
type AlibabaInteractSensorPopwindowAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_interact_sensor_popwindow_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result=0 表示成功
	Result string `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaInteractSensorPopwindowAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = ""
}

var poolAlibabaInteractSensorPopwindowAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaInteractSensorPopwindowAPIResponse)
	},
}

// GetAlibabaInteractSensorPopwindowAPIResponse 从 sync.Pool 获取 AlibabaInteractSensorPopwindowAPIResponse
func GetAlibabaInteractSensorPopwindowAPIResponse() *AlibabaInteractSensorPopwindowAPIResponse {
	return poolAlibabaInteractSensorPopwindowAPIResponse.Get().(*AlibabaInteractSensorPopwindowAPIResponse)
}

// ReleaseAlibabaInteractSensorPopwindowAPIResponse 将 AlibabaInteractSensorPopwindowAPIResponse 保存到 sync.Pool
func ReleaseAlibabaInteractSensorPopwindowAPIResponse(v *AlibabaInteractSensorPopwindowAPIResponse) {
	v.Reset()
	poolAlibabaInteractSensorPopwindowAPIResponse.Put(v)
}
