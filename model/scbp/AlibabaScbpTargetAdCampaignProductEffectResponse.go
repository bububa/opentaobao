package scbp

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
定向推广-获取计划中产品推广效果 APIResponse
alibaba.scbp.target.ad.campaign.product.effect

定向推广-获取计划中产品推广效果
*/
type AlibabaScbpTargetAdCampaignProductEffectAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_scbp_target_ad_campaign_product_effect_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 产品数据
    
    ProductList   []TopP4pQuickProductEffectView `json:"product_list,omitempty" xml:"