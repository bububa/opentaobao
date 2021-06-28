package icbu

import (
    "github.com/bububa/opentaobao/model"
)

/* 
层级属性的子属性获取 APIResponse
alibaba.icbu.category.level.attr.get

用于获取层级属性（车型库）的子属性和属性值
*/
type AlibabaIcbuCategoryLevelAttrGetAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaIcbuCategoryLevelAttrGetResponse `json:"alibaba_icbu_category_level_attr_get_response,omitempty"` 
    AlibabaIcbuCategoryLevelAttrGetResponse
}

/* model for simplify = false
type AlibabaIcbuCategoryLevelAttrGetResponse struct {

    // 返回值
    
    ResultList  *struct {
        AlibabaIcbuCategoryLevelAttrGetResult  *AlibabaIcbuCategoryLevelAttrGetResult `json:"alibaba_icbu_category_level_attr_get_result,omitempty"`
    } `json:"result_list,omitempty"`
    

}
*/

type AlibabaIcbuCategoryLevelAttrGetResponse struct {

    // 返回值
    ResultList   *AlibabaIcbuCategoryLevelAttrGetResult `json:"result_list,omitempty"`

}
