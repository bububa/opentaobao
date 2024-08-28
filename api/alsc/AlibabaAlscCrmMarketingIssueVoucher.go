package alsc

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// AlibabaAlscCrmMarketingIssueVoucher 发券
// alibaba.alsc.crm.marketing.issue.voucher
//
// 提供发券功能
func AlibabaAlscCrmMarketingIssueVoucher(ctx context.Context, clt *core.SDKClient, req *alsc.AlibabaAlscCrmMarketingIssueVoucherAPIRequest, resp *alsc.AlibabaAlscCrmMarketingIssueVoucherAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
