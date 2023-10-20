package interact

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaInteractSensorShareAPIResponse 分享 API返回值
// alibaba.interact.sensor.share
//
// 客户端分享
type AlibabaInteractSensorShareAPIResponse struct {
	model.CommonResponse
	AlibabaInteractSensorShareAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaInteractSensorShareAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaInteractSensorShareAPIResponseModel).Reset()
}

// AlibabaInteractSensorShareAPIResponseModel is 分享 成功返回结果
type AlibabaInteractSensorShareAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_interact_sensor_share_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// return=0表示成功
	Result string `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaInteractSensorShareAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = ""
}

var poolAlibabaInteractSensorShareAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaInteractSensorShareAPIResponse)
	},
}

// GetAlibabaInteractSensorShareAPIResponse 从 sync.Pool 获取 AlibabaInteractSensorShareAPIResponse
func GetAlibabaInteractSensorShareAPIResponse() *AlibabaInteractSensorShareAPIResponse {
	return poolAlibabaInteractSensorShareAPIResponse.Get().(*AlibabaInteractSensorShareAPIResponse)
}

// ReleaseAlibabaInteractSensorShareAPIResponse 将 AlibabaInteractSensorShareAPIResponse 保存到 sync.Pool
func ReleaseAlibabaInteractSensorShareAPIResponse(v *AlibabaInteractSensorShareAPIResponse) {
	v.Reset()
	poolAlibabaInteractSensorShareAPIResponse.Put(v)
}
