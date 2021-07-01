package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaScbpAdCampaignFindCampaignPageAPIRequest
分页查询计划 API请求
alibaba.scbp.ad.campaign.find.campaign.page

分页查询计划 */
type AlibabaScbpAdCampaignFindCampaignPageAPIRequest struct {
	model.Params
	// 请求实体类
	_campaignQuery *CampaignQueryDto
	// 用户信息
	_topContext *TopContextDto
}

// New
