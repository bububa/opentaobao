package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaScbpAdKeywordDeleteKeywordBatchAPIRequest
删除关键词 API请求
alibaba.scbp.ad.keyword.delete.keyword.batch

删除关键词 */
type AlibabaScbpAdKeywordDeleteKeywordBatchAPIRequest struct {
	model.Params
	// 计划id
	_campaignId int64
	// 请求参数
	_keywordQuery *KeywordQuery
	// 用户信息
	_topContext *TopContextDto
}

// New
