package util

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliexpressSocialLocaleGetAPIResponse Locale获取接口 API返回值
// aliexpress.social.locale.get
//
// 新增Locale获取接口
type AliexpressSocialLocaleGetAPIResponse struct {
	model.CommonResponse
	AliexpressSocialLocaleGetAPIResponseModel
}

// Reset 清空结构体
func (m *AliexpressSocialLocaleGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AliexpressSocialLocaleGetAPIResponseModel).Reset()
}

// AliexpressSocialLocaleGetAPIResponseModel is Locale获取接口 成功返回结果
type AliexpressSocialLocaleGetAPIResponseModel struct {
	XMLName xml.Name `xml:"aliexpress_social_locale_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 包类型
	Result *ItemPickPagingResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AliexpressSocialLocaleGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAliexpressSocialLocaleGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AliexpressSocialLocaleGetAPIResponse)
	},
}

// GetAliexpressSocialLocaleGetAPIResponse 从 sync.Pool 获取 AliexpressSocialLocaleGetAPIResponse
func GetAliexpressSocialLocaleGetAPIResponse() *AliexpressSocialLocaleGetAPIResponse {
	return poolAliexpressSocialLocaleGetAPIResponse.Get().(*AliexpressSocialLocaleGetAPIResponse)
}

// ReleaseAliexpressSocialLocaleGetAPIResponse 将 AliexpressSocialLocaleGetAPIResponse 保存到 sync.Pool
func ReleaseAliexpressSocialLocaleGetAPIResponse(v *AliexpressSocialLocaleGetAPIResponse) {
	v.Reset()
	poolAliexpressSocialLocaleGetAPIResponse.Put(v)
}
