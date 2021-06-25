package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
商家会员数据上传 APIResponse
alibaba.tcls.aelophy.merchant.user.upload

商家会员数据上传
*/
type AlibabaTclsAelophyMerchantUserUploadAPIResponse struct {
    model.CommonResponse
    Response *AlibabaTclsAelophyMerchantUserUploadResponse `json:"alibaba_tcls_aelophy_merchant_user_upload_response,omitempty"`
}

type AlibabaTclsAelophyMerchantUserUploadResponse struct {

    // 接口返回model
    ApiResult   *AlibabaTclsAelophyMerchantUserUploadApiResult `json:"api_result,omitempty"`

}
