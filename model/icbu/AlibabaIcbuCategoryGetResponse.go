package icbu

import (
    "github.com/bububa/opentaobao/model"
)

/* 
商品发布类目获取 APIResponse
alibaba.icbu.category.get

获取商品发布类目
*/
type AlibabaIcbuCategoryGetAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaIcbuCategoryGetResponse `json:"alibaba_icbu_category_get_response,omitempty"` 
    AlibabaIcbuCategoryGetResponse
}

/* model for simplify = false
type AlibabaIcbuCategoryGetResponse struct {

    // 类目信息
    
    Category  *struct {
        PostCategory  *PostCategory `json:"post_category,omitempty"`
    } `json:"category,omitempty"`
    

}
*/

type AlibabaIcbuCategoryGetResponse struct {

    // 类目信息
    Category   *PostCategory `json:"category,omitempty"`

}
