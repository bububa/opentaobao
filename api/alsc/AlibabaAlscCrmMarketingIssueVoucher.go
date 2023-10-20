package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// AlibabaAlscCrmMarketingIssueVoucher 发券
// alibaba.alsc.crm.marketing.issue.voucher
//
// 提供发券功能
func AlibabaAlscCrmMarketingIssueVoucher(clt *core.SDKClient, req *alsc.AlibabaAlscCrmMarketingIssueVoucherAPIRequest, resp *alsc.AlibabaAlscCrmMarketingIssueVoucherAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
