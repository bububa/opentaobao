package icbu

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIcbuCategoryAttrvalueGetAPIResponse 属性值获取 API返回值
// alibaba.icbu.category.attrvalue.get
//
// 属性值获取
type AlibabaIcbuCategoryAttrvalueGetAPIResponse struct {
	model.CommonResponse
	AlibabaIcbuCategoryAttrvalueGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaIcbuCategoryAttrvalueGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaIcbuCategoryAttrvalueGetAPIResponseModel).Reset()
}

// AlibabaIcbuCategoryAttrvalueGetAPIResponseModel is 属性值获取 成功返回结果
type AlibabaIcbuCategoryAttrvalueGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_icbu_category_attrvalue_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值
	ResultList []AttributeValue `json:"result_list,omitempty" xml:"result_list>attribute_value,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaIcbuCategoryAttrvalueGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ResultList = m.ResultList[:0]
}

var poolAlibabaIcbuCategoryAttrvalueGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaIcbuCategoryAttrvalueGetAPIResponse)
	},
}

// GetAlibabaIcbuCategoryAttrvalueGetAPIResponse 从 sync.Pool 获取 AlibabaIcbuCategoryAttrvalueGetAPIResponse
func GetAlibabaIcbuCategoryAttrvalueGetAPIResponse() *AlibabaIcbuCategoryAttrvalueGetAPIResponse {
	return poolAlibabaIcbuCategoryAttrvalueGetAPIResponse.Get().(*AlibabaIcbuCategoryAttrvalueGetAPIResponse)
}

// ReleaseAlibabaIcbuCategoryAttrvalueGetAPIResponse 将 AlibabaIcbuCategoryAttrvalueGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaIcbuCategoryAttrvalueGetAPIResponse(v *AlibabaIcbuCategoryAttrvalueGetAPIResponse) {
	v.Reset()
	poolAlibabaIcbuCategoryAttrvalueGetAPIResponse.Put(v)
}
