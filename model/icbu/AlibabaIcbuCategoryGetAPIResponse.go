package icbu

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIcbuCategoryGetAPIResponse 商品发布类目获取 API返回值
// alibaba.icbu.category.get
//
// 获取商品发布类目
type AlibabaIcbuCategoryGetAPIResponse struct {
	model.CommonResponse
	AlibabaIcbuCategoryGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaIcbuCategoryGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaIcbuCategoryGetAPIResponseModel).Reset()
}

// AlibabaIcbuCategoryGetAPIResponseModel is 商品发布类目获取 成功返回结果
type AlibabaIcbuCategoryGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_icbu_category_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 类目信息
	Category *PostCategory `json:"category,omitempty" xml:"category,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaIcbuCategoryGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Category = nil
}

var poolAlibabaIcbuCategoryGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaIcbuCategoryGetAPIResponse)
	},
}

// GetAlibabaIcbuCategoryGetAPIResponse 从 sync.Pool 获取 AlibabaIcbuCategoryGetAPIResponse
func GetAlibabaIcbuCategoryGetAPIResponse() *AlibabaIcbuCategoryGetAPIResponse {
	return poolAlibabaIcbuCategoryGetAPIResponse.Get().(*AlibabaIcbuCategoryGetAPIResponse)
}

// ReleaseAlibabaIcbuCategoryGetAPIResponse 将 AlibabaIcbuCategoryGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaIcbuCategoryGetAPIResponse(v *AlibabaIcbuCategoryGetAPIResponse) {
	v.Reset()
	poolAlibabaIcbuCategoryGetAPIResponse.Put(v)
}
