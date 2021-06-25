package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
品牌查询接口 APIResponse
alibaba.wdk.merchant.brand.query

三江erp对接时，提供品牌查询的接口
*/
type AlibabaWdkMerchantBrandQueryAPIResponse struct {
    model.CommonResponse
    Response *AlibabaWdkMerchantBrandQueryResponse `json:"alibaba_wdk_merchant_brand_query_response,omitempty"`
}

type AlibabaWdkMerchantBrandQueryResponse struct {

    // result
    Result   *AlibabaWdkMerchantBrandQueryResult `json:"result,omitempty"`

}
