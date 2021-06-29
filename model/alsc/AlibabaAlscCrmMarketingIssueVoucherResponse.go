package alsc

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
发券 APIResponse
alibaba.alsc.crm.marketing.issue.voucher

提供发券功能
*/
type AlibabaAlscCrmMarketingIssueVoucherAPIResponse struct {
    model.CommonResponse
    AlibabaAlscCrmMarketingIssueVoucherResponse
}

type AlibabaAlscCrmMarketingIssueVoucherResponse struct {
    XMLName xml.Name `xml:"alibaba_alsc_crm_marketing_issue_voucher_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 接口结果
    
    Result   *CommonResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
