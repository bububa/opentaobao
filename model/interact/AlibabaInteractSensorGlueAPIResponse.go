package interact

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaInteractSensorGlueAPIResponse 视频播放 API返回值
// alibaba.interact.sensor.glue
//
// 视频播放
type AlibabaInteractSensorGlueAPIResponse struct {
	model.CommonResponse
	AlibabaInteractSensorGlueAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaInteractSensorGlueAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaInteractSensorGlueAPIResponseModel).Reset()
}

// AlibabaInteractSensorGlueAPIResponseModel is 视频播放 成功返回结果
type AlibabaInteractSensorGlueAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_interact_sensor_glue_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result=0
	Result string `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaInteractSensorGlueAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = ""
}

var poolAlibabaInteractSensorGlueAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaInteractSensorGlueAPIResponse)
	},
}

// GetAlibabaInteractSensorGlueAPIResponse 从 sync.Pool 获取 AlibabaInteractSensorGlueAPIResponse
func GetAlibabaInteractSensorGlueAPIResponse() *AlibabaInteractSensorGlueAPIResponse {
	return poolAlibabaInteractSensorGlueAPIResponse.Get().(*AlibabaInteractSensorGlueAPIResponse)
}

// ReleaseAlibabaInteractSensorGlueAPIResponse 将 AlibabaInteractSensorGlueAPIResponse 保存到 sync.Pool
func ReleaseAlibabaInteractSensorGlueAPIResponse(v *AlibabaInteractSensorGlueAPIResponse) {
	v.Reset()
	poolAlibabaInteractSensorGlueAPIResponse.Put(v)
}
