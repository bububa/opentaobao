package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaScbpAdKeywordGetAPIRequest
外贸直通车查询关键词 API请求
alibaba.scbp.ad.keyword.get

外贸直通车查询关键词 */
type AlibabaScbpAdKeywordGetAPIRequest struct {
	model.Params
	// KeywordQuery
	_queryDto *KeywordQuery
}

// New
