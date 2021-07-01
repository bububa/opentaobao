package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaScbpEffectKeywordListAPIRequest
关键词报表 API请求
alibaba.scbp.effect.keyword.list

关键词报表 */
type AlibabaScbpEffectKeywordListAPIRequest struct {
	model.Params
	// IKeywordQuery
	_p4pKeywordReportQuery *IKeywordQuery
}

// New
