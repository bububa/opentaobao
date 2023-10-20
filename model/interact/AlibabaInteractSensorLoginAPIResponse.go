package interact

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaInteractSensorLoginAPIResponse 获取登陆页面 API返回值
// alibaba.interact.sensor.login
//
// 获取登陆页面
type AlibabaInteractSensorLoginAPIResponse struct {
	model.CommonResponse
	AlibabaInteractSensorLoginAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaInteractSensorLoginAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaInteractSensorLoginAPIResponseModel).Reset()
}

// AlibabaInteractSensorLoginAPIResponseModel is 获取登陆页面 成功返回结果
type AlibabaInteractSensorLoginAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_interact_sensor_login_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// return=0表示成功
	Result string `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaInteractSensorLoginAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = ""
}

var poolAlibabaInteractSensorLoginAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaInteractSensorLoginAPIResponse)
	},
}

// GetAlibabaInteractSensorLoginAPIResponse 从 sync.Pool 获取 AlibabaInteractSensorLoginAPIResponse
func GetAlibabaInteractSensorLoginAPIResponse() *AlibabaInteractSensorLoginAPIResponse {
	return poolAlibabaInteractSensorLoginAPIResponse.Get().(*AlibabaInteractSensorLoginAPIResponse)
}

// ReleaseAlibabaInteractSensorLoginAPIResponse 将 AlibabaInteractSensorLoginAPIResponse 保存到 sync.Pool
func ReleaseAlibabaInteractSensorLoginAPIResponse(v *AlibabaInteractSensorLoginAPIResponse) {
	v.Reset()
	poolAlibabaInteractSensorLoginAPIResponse.Put(v)
}
