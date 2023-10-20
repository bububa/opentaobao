package interact

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaInteractSensorGcanvasAPIResponse gcanvas API返回值
// alibaba.interact.sensor.gcanvas
//
// gcanvas 功能
type AlibabaInteractSensorGcanvasAPIResponse struct {
	model.CommonResponse
	AlibabaInteractSensorGcanvasAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaInteractSensorGcanvasAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaInteractSensorGcanvasAPIResponseModel).Reset()
}

// AlibabaInteractSensorGcanvasAPIResponseModel is gcanvas 成功返回结果
type AlibabaInteractSensorGcanvasAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_interact_sensor_gcanvas_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result=0
	Result string `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaInteractSensorGcanvasAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = ""
}

var poolAlibabaInteractSensorGcanvasAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaInteractSensorGcanvasAPIResponse)
	},
}

// GetAlibabaInteractSensorGcanvasAPIResponse 从 sync.Pool 获取 AlibabaInteractSensorGcanvasAPIResponse
func GetAlibabaInteractSensorGcanvasAPIResponse() *AlibabaInteractSensorGcanvasAPIResponse {
	return poolAlibabaInteractSensorGcanvasAPIResponse.Get().(*AlibabaInteractSensorGcanvasAPIResponse)
}

// ReleaseAlibabaInteractSensorGcanvasAPIResponse 将 AlibabaInteractSensorGcanvasAPIResponse 保存到 sync.Pool
func ReleaseAlibabaInteractSensorGcanvasAPIResponse(v *AlibabaInteractSensorGcanvasAPIResponse) {
	v.Reset()
	poolAlibabaInteractSensorGcanvasAPIResponse.Put(v)
}
