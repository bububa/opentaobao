package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaScbpAdCampaignCreateAPIRequest
创建计划 API请求
alibaba.scbp.ad.campaign.create

创建计划 */
type AlibabaScbpAdCampaignCreateAPIRequest struct {
	model.Params
	// 返回数据
	_campaignOperation *CampaignOperationDto
	// 用户信息
	_topContext *TopContextDto
}

// New
