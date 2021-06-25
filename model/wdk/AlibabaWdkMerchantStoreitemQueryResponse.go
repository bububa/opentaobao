package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
门店商品信心查询 APIResponse
alibaba.wdk.merchant.storeitem.query

门店商品信心查询
*/
type AlibabaWdkMerchantStoreitemQueryAPIResponse struct {
    model.CommonResponse
    Response *AlibabaWdkMerchantStoreitemQueryResponse `json:"alibaba_wdk_merchant_storeitem_query_response,omitempty"`
}

type AlibabaWdkMerchantStoreitemQueryResponse struct {

    // result
    Result   *AlibabaWdkMerchantStoreitemQueryResult `json:"result,omitempty"`

}