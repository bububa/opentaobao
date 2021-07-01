package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaScbpAdReportGetTargetReportAPIRequest
定向报告 API请求
alibaba.scbp.ad.report.get.target.report

定向报告 */
type AlibabaScbpAdReportGetTargetReportAPIRequest struct {
	model.Params
	// 请求参数
	_targetReportOperation *TargetReportOperationDto
	// 用户信息
	_topContext *TopContextDto
}

// New
