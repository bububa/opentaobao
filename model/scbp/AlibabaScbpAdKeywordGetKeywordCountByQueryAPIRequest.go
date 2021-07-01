package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaScbpAdKeywordGetKeywordCountByQueryAPIRequest
计划关键词数目 API请求
alibaba.scbp.ad.keyword.get.keyword.count.by.query

计划关键词数目 */
type AlibabaScbpAdKeywordGetKeywordCountByQueryAPIRequest struct {
	model.Params
	// 用户信息
	_topContext *TopContextDto
	// 请求参数
	_campaignKeywordQuery *CampaignKeywordQuery
}

// New
