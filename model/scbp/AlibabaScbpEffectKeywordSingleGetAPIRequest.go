package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaScbpEffectKeywordSingleGetAPIRequest
单个关键词效果报表 API请求
alibaba.scbp.effect.keyword.single.get

单个关键词效果报表 */
type AlibabaScbpEffectKeywordSingleGetAPIRequest struct {
	model.Params
	// IKeywordQuery
	_p4pKeywordReportQuery *IKeywordQuery
}

// New
