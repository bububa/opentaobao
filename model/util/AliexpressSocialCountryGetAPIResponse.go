package util

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliexpressSocialCountryGetAPIResponse 获取国家列表 API返回值
// aliexpress.social.country.get
//
// 获取目前AE支持的国家列表
type AliexpressSocialCountryGetAPIResponse struct {
	model.CommonResponse
	AliexpressSocialCountryGetAPIResponseModel
}

// Reset 清空结构体
func (m *AliexpressSocialCountryGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AliexpressSocialCountryGetAPIResponseModel).Reset()
}

// AliexpressSocialCountryGetAPIResponseModel is 获取国家列表 成功返回结果
type AliexpressSocialCountryGetAPIResponseModel struct {
	XMLName xml.Name `xml:"aliexpress_social_country_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// ItemPickPagingResult
	Result *ItemPickPagingResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AliexpressSocialCountryGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAliexpressSocialCountryGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AliexpressSocialCountryGetAPIResponse)
	},
}

// GetAliexpressSocialCountryGetAPIResponse 从 sync.Pool 获取 AliexpressSocialCountryGetAPIResponse
func GetAliexpressSocialCountryGetAPIResponse() *AliexpressSocialCountryGetAPIResponse {
	return poolAliexpressSocialCountryGetAPIResponse.Get().(*AliexpressSocialCountryGetAPIResponse)
}

// ReleaseAliexpressSocialCountryGetAPIResponse 将 AliexpressSocialCountryGetAPIResponse 保存到 sync.Pool
func ReleaseAliexpressSocialCountryGetAPIResponse(v *AliexpressSocialCountryGetAPIResponse) {
	v.Reset()
	poolAliexpressSocialCountryGetAPIResponse.Put(v)
}
