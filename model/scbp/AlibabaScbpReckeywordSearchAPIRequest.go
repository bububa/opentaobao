package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaScbpReckeywordSearchAPIRequest
推荐词-词推词 API请求
alibaba.scbp.reckeyword.search

推荐词-词推词 */
type AlibabaScbpReckeywordSearchAPIRequest struct {
	model.Params
	// RecKeywordQuery
	_queryDto *RecKeywordQuery
}

// New
