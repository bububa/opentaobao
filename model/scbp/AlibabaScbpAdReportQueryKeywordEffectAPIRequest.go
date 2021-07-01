package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaScbpAdReportQueryKeywordEffectAPIRequest
关键词报告 API请求
alibaba.scbp.ad.report.query.keyword.effect

关键词报告 */
type AlibabaScbpAdReportQueryKeywordEffectAPIRequest struct {
	model.Params
	// 请求参数
	_keywordReportOperation *KeywordReportOperationDto
	// 用户信息
	_topContext *TopContextDto
}

// New
