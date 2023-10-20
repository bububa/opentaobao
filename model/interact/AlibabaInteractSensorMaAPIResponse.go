package interact

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaInteractSensorMaAPIResponse 码相关API API返回值
// alibaba.interact.sensor.ma
//
// 码相关API
type AlibabaInteractSensorMaAPIResponse struct {
	model.CommonResponse
	AlibabaInteractSensorMaAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaInteractSensorMaAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaInteractSensorMaAPIResponseModel).Reset()
}

// AlibabaInteractSensorMaAPIResponseModel is 码相关API 成功返回结果
type AlibabaInteractSensorMaAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_interact_sensor_ma_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result=0
	Result string `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaInteractSensorMaAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = ""
}

var poolAlibabaInteractSensorMaAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaInteractSensorMaAPIResponse)
	},
}

// GetAlibabaInteractSensorMaAPIResponse 从 sync.Pool 获取 AlibabaInteractSensorMaAPIResponse
func GetAlibabaInteractSensorMaAPIResponse() *AlibabaInteractSensorMaAPIResponse {
	return poolAlibabaInteractSensorMaAPIResponse.Get().(*AlibabaInteractSensorMaAPIResponse)
}

// ReleaseAlibabaInteractSensorMaAPIResponse 将 AlibabaInteractSensorMaAPIResponse 保存到 sync.Pool
func ReleaseAlibabaInteractSensorMaAPIResponse(v *AlibabaInteractSensorMaAPIResponse) {
	v.Reset()
	poolAlibabaInteractSensorMaAPIResponse.Put(v)
}
