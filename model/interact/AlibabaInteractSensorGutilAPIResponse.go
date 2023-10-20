package interact

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaInteractSensorGutilAPIResponse canvas工具包 API返回值
// alibaba.interact.sensor.gutil
//
// canvas工具包
type AlibabaInteractSensorGutilAPIResponse struct {
	model.CommonResponse
	AlibabaInteractSensorGutilAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaInteractSensorGutilAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaInteractSensorGutilAPIResponseModel).Reset()
}

// AlibabaInteractSensorGutilAPIResponseModel is canvas工具包 成功返回结果
type AlibabaInteractSensorGutilAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_interact_sensor_gutil_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result=0 表示成功
	Result string `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaInteractSensorGutilAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = ""
}

var poolAlibabaInteractSensorGutilAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaInteractSensorGutilAPIResponse)
	},
}

// GetAlibabaInteractSensorGutilAPIResponse 从 sync.Pool 获取 AlibabaInteractSensorGutilAPIResponse
func GetAlibabaInteractSensorGutilAPIResponse() *AlibabaInteractSensorGutilAPIResponse {
	return poolAlibabaInteractSensorGutilAPIResponse.Get().(*AlibabaInteractSensorGutilAPIResponse)
}

// ReleaseAlibabaInteractSensorGutilAPIResponse 将 AlibabaInteractSensorGutilAPIResponse 保存到 sync.Pool
func ReleaseAlibabaInteractSensorGutilAPIResponse(v *AlibabaInteractSensorGutilAPIResponse) {
	v.Reset()
	poolAlibabaInteractSensorGutilAPIResponse.Put(v)
}
