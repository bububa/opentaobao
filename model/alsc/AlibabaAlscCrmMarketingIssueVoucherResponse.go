package alsc

import (
    "github.com/bububa/opentaobao/model"
)

/* 
发券 APIResponse
alibaba.alsc.crm.marketing.issue.voucher

提供发券功能
*/
type AlibabaAlscCrmMarketingIssueVoucherAPIResponse struct {
    model.CommonResponse
    Response *AlibabaAlscCrmMarketingIssueVoucherResponse `json:"alibaba_alsc_crm_marketing_issue_voucher_response,omitempty"`
}

type AlibabaAlscCrmMarketingIssueVoucherResponse struct {

    // 接口结果
    Result   *CommonResult `json:"result,omitempty"`

}
