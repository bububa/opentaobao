package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaScbpAdCampaignFindCampaignEffectAPIRequest
批量查询计划效果数据 API请求
alibaba.scbp.ad.campaign.find.campaign.effect

批量查询计划效果数据 */
type AlibabaScbpAdCampaignFindCampaignEffectAPIRequest struct {
	model.Params
	// 计划id集合
	_campaignIdList []int64
	// 开始时间
	_beginDate string
	// 结束时间
	_endDate string
	// 用户信息
	_topContext *TopContextDto
}

// New
