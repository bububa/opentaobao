package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaScbpTargetAdCampaignTagEffectAPIRequest
定向推广-获取推广计划定向效果数据 API请求
alibaba.scbp.target.ad.campaign.tag.effect

定向推广-获取推广计划定向效果数据 */
type AlibabaScbpTargetAdCampaignTagEffectAPIRequest struct {
	model.Params
	// 效果数据
	_topP4pQuickEffectQuery *TopP4pQuickEffectQuery
}

// New
