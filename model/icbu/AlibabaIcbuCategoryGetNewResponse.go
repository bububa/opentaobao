package icbu

import (
    "github.com/bububa/opentaobao/model"
)

/* 
（新）ICBU类目树获取接口 APIResponse
alibaba.icbu.category.get.new

获取商品发布类目
*/
type AlibabaIcbuCategoryGetNewAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaIcbuCategoryGetNewResponse `json:"alibaba_icbu_category_get_new_response,omitempty"` 
    AlibabaIcbuCategoryGetNewResponse
}

/* model for simplify = false
type AlibabaIcbuCategoryGetNewResponse struct {

    // 类目信息
    
    Category  *struct {
        PostCategory  *PostCategory `json:"post_category,omitempty"`
    } `json:"category,omitempty"`
    

}
*/

type AlibabaIcbuCategoryGetNewResponse struct {

    // 类目信息
    Category   *PostCategory `json:"category,omitempty"`

}
