package icbu

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
类目属性获取 APIResponse
alibaba.icbu.category.attribute.get

根据类目ID获取系统定义的属性
*/
type AlibabaIcbuCategoryAttributeGetAPIResponse struct {
    model.CommonResponse
    AlibabaIcbuCategoryAttributeGetResponse
}

type AlibabaIcbuCategoryAttributeGetResponse struct {
    XMLName xml.Name `xml:"alibaba_icbu_category_attribute_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 类目下的属性和属性值信息
    
    Attributes   []Attribute `json:"attributes,omitempty" xml:"attributes>attribute,omitempty"`
    
    
}
