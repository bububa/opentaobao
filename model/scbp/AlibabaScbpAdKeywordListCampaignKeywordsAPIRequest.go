package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaScbpAdKeywordListCampaignKeywordsAPIRequest
获取计划关键词 API请求
alibaba.scbp.ad.keyword.list.campaign.keywords

获取计划关键词 */
type AlibabaScbpAdKeywordListCampaignKeywordsAPIRequest struct {
	model.Params
	// 计划id
	_campaignId int64
	// 搜索条件
	_campaignKeywordQuery *CampaignKeywordQuery
	// 用户信息
	_topContext *TopContextDto
}

// New
