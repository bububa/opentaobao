package scbp

import (
    "github.com/bububa/opentaobao/model"
)

/* 
定向推广-获取计划中产品推广效果 APIResponse
alibaba.scbp.target.ad.campaign.product.effect

定向推广-获取计划中产品推广效果
*/
type AlibabaScbpTargetAdCampaignProductEffectAPIResponse struct {
    model.CommonResponse
    Response *AlibabaScbpTargetAdCampaignProductEffectResponse `json:"alibaba_scbp_target_ad_campaign_product_effect_response,omitempty"`
}

type AlibabaScbpTargetAdCampaignProductEffectResponse struct {

    // 产品数据
    ProductList   []TopP4pQuickProductEffectView `json:"product_list,omitempty"`

    // 总页数
    TotalPage   int64 `json:"total_page,omitempty"`

    // 总个数
    TotalNum   int64 `json:"total_num,omitempty"`

}
