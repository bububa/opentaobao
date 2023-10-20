package interact

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaInteractSensorGravityAPIResponse 重力感应 API返回值
// alibaba.interact.sensor.gravity
//
// native获取重力感应
type AlibabaInteractSensorGravityAPIResponse struct {
	model.CommonResponse
	AlibabaInteractSensorGravityAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaInteractSensorGravityAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaInteractSensorGravityAPIResponseModel).Reset()
}

// AlibabaInteractSensorGravityAPIResponseModel is 重力感应 成功返回结果
type AlibabaInteractSensorGravityAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_interact_sensor_gravity_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 0表示成功呢
	Result string `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaInteractSensorGravityAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = ""
}

var poolAlibabaInteractSensorGravityAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaInteractSensorGravityAPIResponse)
	},
}

// GetAlibabaInteractSensorGravityAPIResponse 从 sync.Pool 获取 AlibabaInteractSensorGravityAPIResponse
func GetAlibabaInteractSensorGravityAPIResponse() *AlibabaInteractSensorGravityAPIResponse {
	return poolAlibabaInteractSensorGravityAPIResponse.Get().(*AlibabaInteractSensorGravityAPIResponse)
}

// ReleaseAlibabaInteractSensorGravityAPIResponse 将 AlibabaInteractSensorGravityAPIResponse 保存到 sync.Pool
func ReleaseAlibabaInteractSensorGravityAPIResponse(v *AlibabaInteractSensorGravityAPIResponse) {
	v.Reset()
	poolAlibabaInteractSensorGravityAPIResponse.Put(v)
}
