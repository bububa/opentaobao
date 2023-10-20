package interact

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaInteractSensorAuthorizeAPIResponse 客户端授权页 API返回值
// alibaba.interact.sensor.authorize
//
// 客户端授权页
type AlibabaInteractSensorAuthorizeAPIResponse struct {
	model.CommonResponse
	AlibabaInteractSensorAuthorizeAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaInteractSensorAuthorizeAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaInteractSensorAuthorizeAPIResponseModel).Reset()
}

// AlibabaInteractSensorAuthorizeAPIResponseModel is 客户端授权页 成功返回结果
type AlibabaInteractSensorAuthorizeAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_interact_sensor_authorize_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// return=0 表示成功
	Result string `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaInteractSensorAuthorizeAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = ""
}

var poolAlibabaInteractSensorAuthorizeAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaInteractSensorAuthorizeAPIResponse)
	},
}

// GetAlibabaInteractSensorAuthorizeAPIResponse 从 sync.Pool 获取 AlibabaInteractSensorAuthorizeAPIResponse
func GetAlibabaInteractSensorAuthorizeAPIResponse() *AlibabaInteractSensorAuthorizeAPIResponse {
	return poolAlibabaInteractSensorAuthorizeAPIResponse.Get().(*AlibabaInteractSensorAuthorizeAPIResponse)
}

// ReleaseAlibabaInteractSensorAuthorizeAPIResponse 将 AlibabaInteractSensorAuthorizeAPIResponse 保存到 sync.Pool
func ReleaseAlibabaInteractSensorAuthorizeAPIResponse(v *AlibabaInteractSensorAuthorizeAPIResponse) {
	v.Reset()
	poolAlibabaInteractSensorAuthorizeAPIResponse.Put(v)
}
