package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaScbpAdReportGetAccountReportAPIRequest
账户报告 API请求
alibaba.scbp.ad.report.get.account.report

账户报告 */
type AlibabaScbpAdReportGetAccountReportAPIRequest struct {
	model.Params
	// 请求参数
	_accountReportOperation *AccountReportOperationDto
	// 用户信息
	_topContext *TopContextDto
}

// New
