package interact

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaInteractSensorBlowAPIResponse 吹气 API返回值
// alibaba.interact.sensor.blow
//
// 客户端吹气
type AlibabaInteractSensorBlowAPIResponse struct {
	model.CommonResponse
	AlibabaInteractSensorBlowAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaInteractSensorBlowAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaInteractSensorBlowAPIResponseModel).Reset()
}

// AlibabaInteractSensorBlowAPIResponseModel is 吹气 成功返回结果
type AlibabaInteractSensorBlowAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_interact_sensor_blow_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// return=0 表示成功
	Result string `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaInteractSensorBlowAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = ""
}

var poolAlibabaInteractSensorBlowAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaInteractSensorBlowAPIResponse)
	},
}

// GetAlibabaInteractSensorBlowAPIResponse 从 sync.Pool 获取 AlibabaInteractSensorBlowAPIResponse
func GetAlibabaInteractSensorBlowAPIResponse() *AlibabaInteractSensorBlowAPIResponse {
	return poolAlibabaInteractSensorBlowAPIResponse.Get().(*AlibabaInteractSensorBlowAPIResponse)
}

// ReleaseAlibabaInteractSensorBlowAPIResponse 将 AlibabaInteractSensorBlowAPIResponse 保存到 sync.Pool
func ReleaseAlibabaInteractSensorBlowAPIResponse(v *AlibabaInteractSensorBlowAPIResponse) {
	v.Reset()
	poolAlibabaInteractSensorBlowAPIResponse.Put(v)
}
