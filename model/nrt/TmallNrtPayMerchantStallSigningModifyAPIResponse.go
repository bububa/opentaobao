package nrt

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
三级商户进件修改 API返回值 
tmall.nrt.pay.merchant.stall.signing.modify

三级商户进件修改
*/
type TmallNrtPayMerchantStallSigningModifyAPIResponse struct {
    model.CommonResponse
    TmallNrtPayMerchantStallSigningModifyAPIResponseModel
}

// 三级商户进件修改 成功返回结果
type TmallNrtPayMerchantStallSigningModifyAPIResponseModel struct {
    XMLName xml.Name `xml:"tmall_nrt_pay_merchant_stall_signing_modify_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 接口返回model
    Result   *TmallNrtPayMerchantStallSigningModifyResultDo `json:"result,omitempty" xml:"result,omitempty"`
}
