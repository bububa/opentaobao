package icbu

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIcbuCategoryGetNewAPIResponse （新）ICBU类目树获取接口 API返回值
// alibaba.icbu.category.get.new
//
// 获取商品发布类目
type AlibabaIcbuCategoryGetNewAPIResponse struct {
	model.CommonResponse
	AlibabaIcbuCategoryGetNewAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaIcbuCategoryGetNewAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaIcbuCategoryGetNewAPIResponseModel).Reset()
}

// AlibabaIcbuCategoryGetNewAPIResponseModel is （新）ICBU类目树获取接口 成功返回结果
type AlibabaIcbuCategoryGetNewAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_icbu_category_get_new_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 类目信息
	Category *PostCategory `json:"category,omitempty" xml:"category,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaIcbuCategoryGetNewAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Category = nil
}

var poolAlibabaIcbuCategoryGetNewAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaIcbuCategoryGetNewAPIResponse)
	},
}

// GetAlibabaIcbuCategoryGetNewAPIResponse 从 sync.Pool 获取 AlibabaIcbuCategoryGetNewAPIResponse
func GetAlibabaIcbuCategoryGetNewAPIResponse() *AlibabaIcbuCategoryGetNewAPIResponse {
	return poolAlibabaIcbuCategoryGetNewAPIResponse.Get().(*AlibabaIcbuCategoryGetNewAPIResponse)
}

// ReleaseAlibabaIcbuCategoryGetNewAPIResponse 将 AlibabaIcbuCategoryGetNewAPIResponse 保存到 sync.Pool
func ReleaseAlibabaIcbuCategoryGetNewAPIResponse(v *AlibabaIcbuCategoryGetNewAPIResponse) {
	v.Reset()
	poolAlibabaIcbuCategoryGetNewAPIResponse.Put(v)
}
