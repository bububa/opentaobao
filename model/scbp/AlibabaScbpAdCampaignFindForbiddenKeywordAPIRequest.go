package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaScbpAdCampaignFindForbiddenKeywordAPIRequest
查询屏蔽词 API请求
alibaba.scbp.ad.campaign.find.forbidden.keyword

查询屏蔽词 */
type AlibabaScbpAdCampaignFindForbiddenKeywordAPIRequest struct {
	model.Params
	// 计划id
	_campaignId int64
	// 用户信息
	_topContext *TopContextDto
}

// New
