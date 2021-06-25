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
    Response *AlibabaWholesaleCategoryGetResponse `json:"alibaba_wholesale_category_get_response,omitempty"`
}

type AlibabaWholesaleCategoryGetResponse struct {

    // 类目结果
    WholesaleCategoryResult   *WholesaleCategoryOpenResult `json:"wholesale_category_result,omitempty"`

}
