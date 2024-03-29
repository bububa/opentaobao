package icbudropshipping

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDropshippingTokenCreateAPIResponse 国际站dropshipping 选品token 创建 API返回值
// alibaba.dropshipping.token.create
//
// 国际站dropshipping 选品token 创建，用于让买家有权限访问我们指定的 商品场馆
type AlibabaDropshippingTokenCreateAPIResponse struct {
	model.CommonResponse
	AlibabaDropshippingTokenCreateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaDropshippingTokenCreateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaDropshippingTokenCreateAPIResponseModel).Reset()
}

// AlibabaDropshippingTokenCreateAPIResponseModel is 国际站dropshipping 选品token 创建 成功返回结果
type AlibabaDropshippingTokenCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_dropshipping_token_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// ecology_token
	EcologyToken string `json:"ecology_token,omitempty" xml:"ecology_token,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaDropshippingTokenCreateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.EcologyToken = ""
}

var poolAlibabaDropshippingTokenCreateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaDropshippingTokenCreateAPIResponse)
	},
}

// GetAlibabaDropshippingTokenCreateAPIResponse 从 sync.Pool 获取 AlibabaDropshippingTokenCreateAPIResponse
func GetAlibabaDropshippingTokenCreateAPIResponse() *AlibabaDropshippingTokenCreateAPIResponse {
	return poolAlibabaDropshippingTokenCreateAPIResponse.Get().(*AlibabaDropshippingTokenCreateAPIResponse)
}

// ReleaseAlibabaDropshippingTokenCreateAPIResponse 将 AlibabaDropshippingTokenCreateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaDropshippingTokenCreateAPIResponse(v *AlibabaDropshippingTokenCreateAPIResponse) {
	v.Reset()
	poolAlibabaDropshippingTokenCreateAPIResponse.Put(v)
}
