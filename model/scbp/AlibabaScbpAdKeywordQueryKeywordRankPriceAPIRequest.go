package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaScbpAdKeywordQueryKeywordRankPriceAPIRequest
查询关键词前五名排价 API请求
alibaba.scbp.ad.keyword.query.keyword.rank.price

查询关键词前五名排价 */
type AlibabaScbpAdKeywordQueryKeywordRankPriceAPIRequest struct {
	model.Params
	// 计划id
	_campaignId int64
	// 请求参数
	_keywordQuery *KeywordQuery
	// 用户信息
	_topContext *TopContextDto
}

// New
