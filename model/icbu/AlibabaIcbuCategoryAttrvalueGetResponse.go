package icbu

import (
    "github.com/bububa/opentaobao/model"
)

/* 
属性值获取 APIResponse
alibaba.icbu.category.attrvalue.get

属性值获取
*/
type AlibabaIcbuCategoryAttrvalueGetAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaIcbuCategoryAttrvalueGetResponse `json:"alibaba_icbu_category_attrvalue_get_response,omitempty"` 
    AlibabaIcbuCategoryAttrvalueGetResponse
}

/* model for simplify = false
type AlibabaIcbuCategoryAttrvalueGetResponse struct {

    // 返回值
    
    ResultList  struct {
        AttributeValue  []AttributeValue `json:"attribute_value,omitempty"`
    } `json:"result_list,omitempty"`
    

}
*/

type AlibabaIcbuCategoryAttrvalueGetResponse struct {

    // 返回值
    ResultList   []AttributeValue `json:"result_list,omitempty"`

}
