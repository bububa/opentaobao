package icbu

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
属性值获取 APIResponse
alibaba.icbu.category.attrvalue.get

属性值获取
*/
type AlibabaIcbuCategoryAttrvalueGetAPIResponse struct {
    model.CommonResponse
    AlibabaIcbuCategoryAttrvalueGetResponse
}

type AlibabaIcbuCategoryAttrvalueGetResponse struct {
    XMLName xml.Name `xml:"alibaba_icbu_category_attrvalue_get_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回值
    
    ResultList   []AttributeValue `json:"result_list,omitempty" xml:"result_list>attribute_value,omitempty"`
    
    
}
