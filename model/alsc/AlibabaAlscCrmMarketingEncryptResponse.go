package alsc

import (
    "github.com/bububa/opentaobao/model"
)

/* 
加密 APIResponse
alibaba.alsc.crm.marketing.encrypt

加密
*/
type AlibabaAlscCrmMarketingEncryptAPIResponse struct {
    model.CommonResponse
    Response *AlibabaAlscCrmMarketingEncryptResponse `json:"alibaba_alsc_crm_marketing_encrypt_response,omitempty"`
}

type AlibabaAlscCrmMarketingEncryptResponse struct {

    // 返回结果
    Result   *CommonResult `json:"result,omitempty"`

}