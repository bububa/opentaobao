package category

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取类目信息 APIResponse
alibaba.wholesale.category.get

获取类目信息
*/
type AlibabaWholesaleCategoryGetAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaWholesaleCategoryGetResponse `json:"alibaba_wholesale_category_get_response,omitempty"` 
    AlibabaWholesaleCategoryGetResponse
}

/* model for simplify = false
type AlibabaWholesaleCategoryGetResponse struct {

    // 类目结果
    
    WholesaleCategoryResult  *struct {
        WholesaleCategoryOpenResult  *WholesaleCategoryOpenResult `json:"wholesale_category_open_result,omitempty"`
    } `json:"wholesale_category_result,omitempty"`
    

}
*/

type AlibabaWholesaleCategoryGetResponse struct {

    // 类目结果
    WholesaleCategoryResult   *WholesaleCategoryOpenResult `json:"wholesale_category_result,omitempty"`

}
