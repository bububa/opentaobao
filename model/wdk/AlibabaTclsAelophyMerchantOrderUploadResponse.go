package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
商家订单数据上传 APIResponse
alibaba.tcls.aelophy.merchant.order.upload

商家订单数据上传
*/
type AlibabaTclsAelophyMerchantOrderUploadAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaTclsAelophyMerchantOrderUploadResponse `json:"alibaba_tcls_aelophy_merchant_order_upload_response,omitempty"` 
    AlibabaTclsAelophyMerchantOrderUploadResponse
}

/* model for simplify = false
type AlibabaTclsAelophyMerchantOrderUploadResponse struct {

    // 根据站点名称查询产品
    
    ApiResult  *struct {
        AlibabaTclsAelophyMerchantOrderUploadApiResult  *AlibabaTclsAelophyMerchantOrderUploadApiResult `json:"alibaba_tcls_aelophy_merchant_order_upload_api_result,omitempty"`
    } `json:"api_result,omitempty"`
    

}
*/

type AlibabaTclsAelophyMerchantOrderUploadResponse struct {

    // 根据站点名称查询产品
    ApiResult   *AlibabaTclsAelophyMerchantOrderUploadApiResult `json:"api_result,omitempty"`

}
