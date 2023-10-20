package interact

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaInteractSensorGmediaAPIResponse gmedia API返回值
// alibaba.interact.sensor.gmedia
//
// 媒体功能
type AlibabaInteractSensorGmediaAPIResponse struct {
	model.CommonResponse
	AlibabaInteractSensorGmediaAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaInteractSensorGmediaAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaInteractSensorGmediaAPIResponseModel).Reset()
}

// AlibabaInteractSensorGmediaAPIResponseModel is gmedia 成功返回结果
type AlibabaInteractSensorGmediaAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_interact_sensor_gmedia_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result=0 表示成功
	Result string `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaInteractSensorGmediaAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = ""
}

var poolAlibabaInteractSensorGmediaAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaInteractSensorGmediaAPIResponse)
	},
}

// GetAlibabaInteractSensorGmediaAPIResponse 从 sync.Pool 获取 AlibabaInteractSensorGmediaAPIResponse
func GetAlibabaInteractSensorGmediaAPIResponse() *AlibabaInteractSensorGmediaAPIResponse {
	return poolAlibabaInteractSensorGmediaAPIResponse.Get().(*AlibabaInteractSensorGmediaAPIResponse)
}

// ReleaseAlibabaInteractSensorGmediaAPIResponse 将 AlibabaInteractSensorGmediaAPIResponse 保存到 sync.Pool
func ReleaseAlibabaInteractSensorGmediaAPIResponse(v *AlibabaInteractSensorGmediaAPIResponse) {
	v.Reset()
	poolAlibabaInteractSensorGmediaAPIResponse.Put(v)
}
