package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
商家会员数据上传 API返回值 
alibaba.tcls.aelophy.merchant.user.upload

商家会员数据上传
*/
type AlibabaTclsAelophyMerchantUserUploadAPIResponse struct {
    model.CommonResponse
    AlibabaTclsAelophyMerchantUserUploadResponse
}

// 商家会员数据上传 成功返回结果
type AlibabaTclsAelophyMerchantUserUploadResponse struct {
    XMLName xml.Name `xml:"alibaba_tcls_aelophy_merchant_user_upload_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 接口返回model
    ApiResult   *AlibabaTclsAelophyMerchantUserUploadApiResult `json:"api_result,omitempty" xml:"api_result,omitempty"`
}
