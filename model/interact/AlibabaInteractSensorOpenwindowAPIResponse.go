package interact

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaInteractSensorOpenwindowAPIResponse 客户端打开新页面 API返回值
// alibaba.interact.sensor.openwindow
//
// 客户端打开新页面
type AlibabaInteractSensorOpenwindowAPIResponse struct {
	model.CommonResponse
	AlibabaInteractSensorOpenwindowAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaInteractSensorOpenwindowAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaInteractSensorOpenwindowAPIResponseModel).Reset()
}

// AlibabaInteractSensorOpenwindowAPIResponseModel is 客户端打开新页面 成功返回结果
type AlibabaInteractSensorOpenwindowAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_interact_sensor_openwindow_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result=0
	Result string `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaInteractSensorOpenwindowAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = ""
}

var poolAlibabaInteractSensorOpenwindowAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaInteractSensorOpenwindowAPIResponse)
	},
}

// GetAlibabaInteractSensorOpenwindowAPIResponse 从 sync.Pool 获取 AlibabaInteractSensorOpenwindowAPIResponse
func GetAlibabaInteractSensorOpenwindowAPIResponse() *AlibabaInteractSensorOpenwindowAPIResponse {
	return poolAlibabaInteractSensorOpenwindowAPIResponse.Get().(*AlibabaInteractSensorOpenwindowAPIResponse)
}

// ReleaseAlibabaInteractSensorOpenwindowAPIResponse 将 AlibabaInteractSensorOpenwindowAPIResponse 保存到 sync.Pool
func ReleaseAlibabaInteractSensorOpenwindowAPIResponse(v *AlibabaInteractSensorOpenwindowAPIResponse) {
	v.Reset()
	poolAlibabaInteractSensorOpenwindowAPIResponse.Put(v)
}
