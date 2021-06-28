package icbu

import (
    "github.com/bububa/opentaobao/model"
)

/* 
类目属性获取 APIResponse
alibaba.icbu.category.attribute.get

根据类目ID获取系统定义的属性
*/
type AlibabaIcbuCategoryAttributeGetAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaIcbuCategoryAttributeGetResponse `json:"alibaba_icbu_category_attribute_get_response,omitempty"` 
    AlibabaIcbuCategoryAttributeGetResponse
}

/* model for simplify = false
type AlibabaIcbuCategoryAttributeGetResponse struct {

    // 类目下的属性和属性值信息
    
    Attributes  struct {
        Attribute  []Attribute `json:"attribute,omitempty"`
    } `json:"attributes,omitempty"`
    

}
*/

type AlibabaIcbuCategoryAttributeGetResponse struct {

    // 类目下的属性和属性值信息
    Attributes   []Attribute `json:"attributes,omitempty"`

}
