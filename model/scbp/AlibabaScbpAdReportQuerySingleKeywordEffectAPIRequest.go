package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaScbpAdReportQuerySingleKeywordEffectAPIRequest
单个关键词报告 API请求
alibaba.scbp.ad.report.query.single.keyword.effect

单个关键词报告 */
type AlibabaScbpAdReportQuerySingleKeywordEffectAPIRequest struct {
	model.Params
	// 返回详情
	_keywordReportOperation *KeywordReportOperationDto
	// 用户信息
	_topContext *TopContextDto
}

// New
