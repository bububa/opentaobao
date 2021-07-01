package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaScbpAdKeywordUpdateKeywordPriceBatchAPIRequest
修改关键词价格 API请求
alibaba.scbp.ad.keyword.update.keyword.price.batch

修改关键词价格 */
type AlibabaScbpAdKeywordUpdateKeywordPriceBatchAPIRequest struct {
	model.Params
	// 计划id
	_campaignId int64
	// 请求参数
	_keywordUpdateQuery *KeywordUpdateQuery
	// 用户信息
	_topContext *TopContextDto
}

// New
