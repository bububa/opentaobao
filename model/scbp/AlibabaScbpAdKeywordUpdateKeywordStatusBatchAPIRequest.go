package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaScbpAdKeywordUpdateKeywordStatusBatchAPIRequest
修改关键词状态 API请求
alibaba.scbp.ad.keyword.update.keyword.status.batch

修改关键词状态 */
type AlibabaScbpAdKeywordUpdateKeywordStatusBatchAPIRequest struct {
	model.Params
	// 计划id
	_campaignId int64
	// 更新数据
	_keywordUpdateQuery *KeywordUpdateQuery
	// 用户信息
	_topContext *TopContextDto
}

// New
