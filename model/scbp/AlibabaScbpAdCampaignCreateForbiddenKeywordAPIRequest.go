package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaScbpAdCampaignCreateForbiddenKeywordAPIRequest
创建屏蔽词 API请求
alibaba.scbp.ad.campaign.create.forbidden.keyword

创建屏蔽词 */
type AlibabaScbpAdCampaignCreateForbiddenKeywordAPIRequest struct {
	model.Params
	// 请求参数
	_forbiddenKeywordBatchOperation *ForbiddenKeywordBatchOperationDto
	// 计划id
	_campaignId int64
	// 用户信息
	_topContext *TopContextDto
}

// New
