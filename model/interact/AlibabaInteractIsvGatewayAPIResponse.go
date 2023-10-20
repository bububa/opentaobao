package interact

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaInteractIsvGatewayAPIResponse isv调用gateway API返回值
// alibaba.interact.isv.gateway
//
// isv能够调用jae本身的server
type AlibabaInteractIsvGatewayAPIResponse struct {
	model.CommonResponse
	AlibabaInteractIsvGatewayAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaInteractIsvGatewayAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaInteractIsvGatewayAPIResponseModel).Reset()
}

// AlibabaInteractIsvGatewayAPIResponseModel is isv调用gateway 成功返回结果
type AlibabaInteractIsvGatewayAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_interact_isv_gateway_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// ret=0
	Ret string `json:"ret,omitempty" xml:"ret,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaInteractIsvGatewayAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Ret = ""
}

var poolAlibabaInteractIsvGatewayAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaInteractIsvGatewayAPIResponse)
	},
}

// GetAlibabaInteractIsvGatewayAPIResponse 从 sync.Pool 获取 AlibabaInteractIsvGatewayAPIResponse
func GetAlibabaInteractIsvGatewayAPIResponse() *AlibabaInteractIsvGatewayAPIResponse {
	return poolAlibabaInteractIsvGatewayAPIResponse.Get().(*AlibabaInteractIsvGatewayAPIResponse)
}

// ReleaseAlibabaInteractIsvGatewayAPIResponse 将 AlibabaInteractIsvGatewayAPIResponse 保存到 sync.Pool
func ReleaseAlibabaInteractIsvGatewayAPIResponse(v *AlibabaInteractIsvGatewayAPIResponse) {
	v.Reset()
	poolAlibabaInteractIsvGatewayAPIResponse.Put(v)
}
