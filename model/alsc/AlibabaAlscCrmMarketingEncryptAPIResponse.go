package alsc

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
加密 API返回值 
alibaba.alsc.crm.marketing.encrypt

加密
*/
type AlibabaAlscCrmMarketingEncryptAPIResponse struct {
    model.CommonResponse
    AlibabaAlscCrmMarketingEncryptAPIResponseModel
}

// 加密 成功返回结果
type AlibabaAlscCrmMarketingEncryptAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_alsc_crm_marketing_encrypt_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回结果
    Result   *CommonResult `json:"result,omitempty" xml:"result,omitempty"`
}
