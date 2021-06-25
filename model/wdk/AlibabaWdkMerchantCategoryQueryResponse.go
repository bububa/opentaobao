package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
三江erp对接类目查询接口 APIResponse
alibaba.wdk.merchant.category.query

三江erp对接类目查询接口
*/
type AlibabaWdkMerchantCategoryQueryAPIResponse struct {
    model.CommonResponse
    Response *AlibabaWdkMerchantCategoryQueryResponse `json:"alibaba_wdk_merchant_category_query_response,omitempty"`
}

type AlibabaWdkMerchantCategoryQueryResponse struct {

    // success
    Suc   bool `json:"suc,omitempty"`

    // errorCode
    Errorcode   string `json:"errorcode,omitempty"`

    // errorDesc
    Errordesc   string `json:"errordesc,omitempty"`

    // result
    Result   string `json:"result,omitempty"`

}
