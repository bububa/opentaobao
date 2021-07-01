package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlscCrmMarketingIssueVoucherAPIRequest
发券 API请求
alibaba.alsc.crm.marketing.issue.voucher

提供发券功能 */
type AlibabaAlscCrmMarketingIssueVoucherAPIRequest struct {
	model.Params
	// 参数
	_paramIssueVoucherReq *IssueVoucherReq
}

// New
