package category

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliexpressSocialDiscategoryGetAPIResponse 展示类目获取接口 API返回值
// aliexpress.social.discategory.get
//
// AE展示类目获取接口
type AliexpressSocialDiscategoryGetAPIResponse struct {
	model.CommonResponse
	AliexpressSocialDiscategoryGetAPIResponseModel
}

// Reset 清空结构体
func (m *AliexpressSocialDiscategoryGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AliexpressSocialDiscategoryGetAPIResponseModel).Reset()
}

// AliexpressSocialDiscategoryGetAPIResponseModel is 展示类目获取接口 成功返回结果
type AliexpressSocialDiscategoryGetAPIResponseModel struct {
	XMLName xml.Name `xml:"aliexpress_social_discategory_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *ItemPickPagingResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AliexpressSocialDiscategoryGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAliexpressSocialDiscategoryGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AliexpressSocialDiscategoryGetAPIResponse)
	},
}

// GetAliexpressSocialDiscategoryGetAPIResponse 从 sync.Pool 获取 AliexpressSocialDiscategoryGetAPIResponse
func GetAliexpressSocialDiscategoryGetAPIResponse() *AliexpressSocialDiscategoryGetAPIResponse {
	return poolAliexpressSocialDiscategoryGetAPIResponse.Get().(*AliexpressSocialDiscategoryGetAPIResponse)
}

// ReleaseAliexpressSocialDiscategoryGetAPIResponse 将 AliexpressSocialDiscategoryGetAPIResponse 保存到 sync.Pool
func ReleaseAliexpressSocialDiscategoryGetAPIResponse(v *AliexpressSocialDiscategoryGetAPIResponse) {
	v.Reset()
	poolAliexpressSocialDiscategoryGetAPIResponse.Put(v)
}
