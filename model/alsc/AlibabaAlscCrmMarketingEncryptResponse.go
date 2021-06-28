package alsc

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
加密 APIResponse
alibaba.alsc.crm.marketing.encrypt

加密
*/
type AlibabaAlscCrmMarketingEncryptAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_alsc_crm_marketing_encrypt_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回结果
    
    Result   *CommonResult `json:"result,omitempty" xml:"