package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaScbpAdKeywordRankPriceBatchgetAPIRequest
外贸直通车关键词前五名批量排价 API请求
alibaba.scbp.ad.keyword.rank.price.batchget

外贸直通车关键词前五名批量排价 */
type AlibabaScbpAdKeywordRankPriceBatchgetAPIRequest struct {
	model.Params
	// 上下文
	_context *ContextDto
	// keyword_request
	_keywordRequest *TopKeywordListDto
}

// New
