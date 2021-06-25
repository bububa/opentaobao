package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
修改类目 APIResponse
alibaba.wdk.item.category.update

修改类目
*/
type AlibabaWdkItemCategoryUpdateAPIResponse struct {
    model.CommonResponse
    Response *AlibabaWdkItemCategoryUpdateResponse `json:"alibaba_wdk_item_category_update_response,omitempty"`
}

type AlibabaWdkItemCategoryUpdateResponse struct {

    // result
    Result   *AlibabaWdkItemCategoryUpdateResult `json:"result,omitempty"`

}
