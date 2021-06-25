package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
商家类目新增接口 APIResponse
alibaba.wdk.sku.category.add

商家类目新增接口
*/
type AlibabaWdkSkuCategoryAddAPIResponse struct {
    model.CommonResponse
    Response *AlibabaWdkSkuCategoryAddResponse `json:"alibaba_wdk_sku_category_add_response,omitempty"`
}

type AlibabaWdkSkuCategoryAddResponse struct {

    // 调用结果
    Result   *AlibabaWdkSkuCategoryAddApiResult `json:"result,omitempty"`

}
