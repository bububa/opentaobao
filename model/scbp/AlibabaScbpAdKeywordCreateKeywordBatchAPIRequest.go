package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaScbpAdKeywordCreateKeywordBatchAPIRequest
关键词添加 API请求
alibaba.scbp.ad.keyword.create.keyword.batch

关键词添加 */
type AlibabaScbpAdKeywordCreateKeywordBatchAPIRequest struct {
	model.Params
	// 计划id
	_campaignId int64
	// 请求参数
	_keywordQuery *KeywordQuery
	// 用户信息
	_topContext *TopContextDto
}

// New
