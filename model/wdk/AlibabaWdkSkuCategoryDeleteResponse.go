package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
商家类目删除接口 APIResponse
alibaba.wdk.sku.category.delete

商家类目删除接口
*/
type AlibabaWdkSkuCategoryDeleteAPIResponse struct {
    model.CommonResponse
    Response *AlibabaWdkSkuCategoryDeleteResponse `json:"alibaba_wdk_sku_category_delete_response,omitempty"`
}

type AlibabaWdkSkuCategoryDeleteResponse struct {

    // 调用结果
    Result   *AlibabaWdkSkuCategoryDeleteApiResult `json:"result,omitempty"`

}
