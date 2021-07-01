package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaScbpAdCampaignUpdateAPIRequest
修改计划 API请求
alibaba.scbp.ad.campaign.update

修改计划 */
type AlibabaScbpAdCampaignUpdateAPIRequest struct {
	model.Params
	// 修改数据
	_campaignOperation *CampaignOperationDto
	// 用户信息
	_topContext *TopContextDto
}

// New
