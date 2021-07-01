package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaScbpTargetAdCampaignProductEffectAPIRequest
定向推广-获取计划中产品推广效果 API请求
alibaba.scbp.target.ad.campaign.product.effect

定向推广-获取计划中产品推广效果 */
type AlibabaScbpTargetAdCampaignProductEffectAPIRequest struct {
	model.Params
	// TopP4pQuickEffectQuery
	_topP4pQuickEffectQuery *TopP4pQuickEffectQuery
}

// New
