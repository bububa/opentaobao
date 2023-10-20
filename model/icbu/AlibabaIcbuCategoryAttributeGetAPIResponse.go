package icbu

import (
	"encoding/xml"

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

// AlibabaIcbuCategoryAttributeGetAPIResponseModel is 类目属性获取 成功返回结果
type AlibabaIcbuCategoryAttributeGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_icbu_category_attribute_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 类目下的属性和属性值信息
	Attributes []Attribute `json:"attributes,omitempty" xml:"attributes>attribute,omitempty"`
}
