package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaScbpAdKeywordRankPriceGetAPIRequest
外贸直通车关键词前五名排价 API请求
alibaba.scbp.ad.keyword.rank.price.get

外贸直通车关键词前五名排价 */
type AlibabaScbpAdKeywordRankPriceGetAPIRequest struct {
	model.Params
	// 关键词
	_keyword string
}

// New
