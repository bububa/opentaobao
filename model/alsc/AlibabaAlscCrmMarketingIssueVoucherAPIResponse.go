package alsc

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscCrmMarketingIssueVoucherAPIResponse 发券 API返回值
// alibaba.alsc.crm.marketing.issue.voucher
//
// 提供发券功能
type AlibabaAlscCrmMarketingIssueVoucherAPIResponse struct {
	model.CommonResponse
	AlibabaAlscCrmMarketingIssueVoucherAPIResponseModel
}

// AlibabaAlscCrmMarketingIssueVoucherAPIResponseModel is 发券 成功返回结果
type AlibabaAlscCrmMarketingIssueVoucherAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alsc_crm_marketing_issue_voucher_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口结果
	Result *CommonResult `json:"result,omitempty" xml:"result,omitempty"`
}
