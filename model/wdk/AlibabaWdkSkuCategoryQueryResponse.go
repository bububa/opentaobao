package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
商家类目获取接口 APIResponse
alibaba.wdk.sku.category.query

商家类目获取接口
*/
type AlibabaWdkSkuCategoryQueryAPIResponse struct {
    model.CommonResponse
    Response *AlibabaWdkSkuCategoryQueryResponse `json:"alibaba_wdk_sku_category_query_response,omitempty"`
}

type AlibabaWdkSkuCategoryQueryResponse struct {

    // 调用结果
    Result   *AlibabaWdkSkuCategoryQueryApiResult `json:"result,omitempty"`

}
