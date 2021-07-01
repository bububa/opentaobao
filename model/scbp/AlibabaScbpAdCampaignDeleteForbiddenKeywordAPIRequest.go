package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaScbpAdCampaignDeleteForbiddenKeywordAPIRequest
删除屏蔽词 API请求
alibaba.scbp.ad.campaign.delete.forbidden.keyword

删除屏蔽词 */
type AlibabaScbpAdCampaignDeleteForbiddenKeywordAPIRequest struct {
	model.Params
	// 计划id
	_campaignId int64
	// 请求参数
	_forbiddenKeywordBatchOperation *ForbiddenKeywordBatchOperationDto
	// 用户信息
	_topContext *TopContextDto
}

// New
