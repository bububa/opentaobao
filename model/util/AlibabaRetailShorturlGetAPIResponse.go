package util

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaRetailShorturlGetAPIResponse 短链接获取 API返回值
// alibaba.retail.shorturl.get
//
// 短链接获取
type AlibabaRetailShorturlGetAPIResponse struct {
	model.CommonResponse
	AlibabaRetailShorturlGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaRetailShorturlGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaRetailShorturlGetAPIResponseModel).Reset()
}

// AlibabaRetailShorturlGetAPIResponseModel is 短链接获取 成功返回结果
type AlibabaRetailShorturlGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_retail_shorturl_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaRetailShorturlGetResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaRetailShorturlGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaRetailShorturlGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaRetailShorturlGetAPIResponse)
	},
}

// GetAlibabaRetailShorturlGetAPIResponse 从 sync.Pool 获取 AlibabaRetailShorturlGetAPIResponse
func GetAlibabaRetailShorturlGetAPIResponse() *AlibabaRetailShorturlGetAPIResponse {
	return poolAlibabaRetailShorturlGetAPIResponse.Get().(*AlibabaRetailShorturlGetAPIResponse)
}

// ReleaseAlibabaRetailShorturlGetAPIResponse 将 AlibabaRetailShorturlGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaRetailShorturlGetAPIResponse(v *AlibabaRetailShorturlGetAPIResponse) {
	v.Reset()
	poolAlibabaRetailShorturlGetAPIResponse.Put(v)
}
