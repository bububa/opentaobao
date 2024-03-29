package aetools

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliexpressAffiliateLinkGenerateAPIResponse 联盟推广链接生成 API返回值
// aliexpress.affiliate.link.generate
//
// AE联盟推广链接生成接口
type AliexpressAffiliateLinkGenerateAPIResponse struct {
	model.CommonResponse
	AliexpressAffiliateLinkGenerateAPIResponseModel
}

// Reset 清空结构体
func (m *AliexpressAffiliateLinkGenerateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AliexpressAffiliateLinkGenerateAPIResponseModel).Reset()
}

// AliexpressAffiliateLinkGenerateAPIResponseModel is 联盟推广链接生成 成功返回结果
type AliexpressAffiliateLinkGenerateAPIResponseModel struct {
	XMLName xml.Name `xml:"aliexpress_affiliate_link_generate_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	RespResult *ResponseDto `json:"resp_result,omitempty" xml:"resp_result,omitempty"`
}

// Reset 清空结构体
func (m *AliexpressAffiliateLinkGenerateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.RespResult = nil
}

var poolAliexpressAffiliateLinkGenerateAPIResponse = sync.Pool{
	New: func() any {
		return new(AliexpressAffiliateLinkGenerateAPIResponse)
	},
}

// GetAliexpressAffiliateLinkGenerateAPIResponse 从 sync.Pool 获取 AliexpressAffiliateLinkGenerateAPIResponse
func GetAliexpressAffiliateLinkGenerateAPIResponse() *AliexpressAffiliateLinkGenerateAPIResponse {
	return poolAliexpressAffiliateLinkGenerateAPIResponse.Get().(*AliexpressAffiliateLinkGenerateAPIResponse)
}

// ReleaseAliexpressAffiliateLinkGenerateAPIResponse 将 AliexpressAffiliateLinkGenerateAPIResponse 保存到 sync.Pool
func ReleaseAliexpressAffiliateLinkGenerateAPIResponse(v *AliexpressAffiliateLinkGenerateAPIResponse) {
	v.Reset()
	poolAliexpressAffiliateLinkGenerateAPIResponse.Put(v)
}
