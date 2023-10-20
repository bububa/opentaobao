package icbu

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIcbuCategoryAttributeGetAPIResponse 类目属性获取 API返回值
// alibaba.icbu.category.attribute.get
//
// 根据类目ID获取系统定义的属性
type AlibabaIcbuCategoryAttributeGetAPIResponse struct {
	model.CommonResponse
	AlibabaIcbuCategoryAttributeGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaIcbuCategoryAttributeGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaIcbuCategoryAttributeGetAPIResponseModel).Reset()
}

// AlibabaIcbuCategoryAttributeGetAPIResponseModel is 类目属性获取 成功返回结果
type AlibabaIcbuCategoryAttributeGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_icbu_category_attribute_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 类目下的属性和属性值信息
	Attributes []Attribute `json:"attributes,omitempty" xml:"attributes>attribute,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaIcbuCategoryAttributeGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Attributes = m.Attributes[:0]
}

var poolAlibabaIcbuCategoryAttributeGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaIcbuCategoryAttributeGetAPIResponse)
	},
}

// GetAlibabaIcbuCategoryAttributeGetAPIResponse 从 sync.Pool 获取 AlibabaIcbuCategoryAttributeGetAPIResponse
func GetAlibabaIcbuCategoryAttributeGetAPIResponse() *AlibabaIcbuCategoryAttributeGetAPIResponse {
	return poolAlibabaIcbuCategoryAttributeGetAPIResponse.Get().(*AlibabaIcbuCategoryAttributeGetAPIResponse)
}

// ReleaseAlibabaIcbuCategoryAttributeGetAPIResponse 将 AlibabaIcbuCategoryAttributeGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaIcbuCategoryAttributeGetAPIResponse(v *AlibabaIcbuCategoryAttributeGetAPIResponse) {
	v.Reset()
	poolAlibabaIcbuCategoryAttributeGetAPIResponse.Put(v)
}
