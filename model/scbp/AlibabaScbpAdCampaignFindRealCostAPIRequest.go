package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaScbpAdCampaignFindRealCostAPIRequest
批量查询计划消耗数据 API请求
alibaba.scbp.ad.campaign.find.real.cost

批量查询计划消耗数据 */
type AlibabaScbpAdCampaignFindRealCostAPIRequest struct {
	model.Params
	// 系统自动生成
	_campaignQuery *CampaignQueryDto
	// 用户信息
	_topContext *TopContextDto
}

// New
