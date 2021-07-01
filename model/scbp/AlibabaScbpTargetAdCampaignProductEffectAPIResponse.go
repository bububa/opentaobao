package scbp

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
定向推广-获取计划中产品推广效果 API返回值 
alibaba.scbp.target.ad.campaign.product.effect

定向推广-获取计划中产品推广效果
*/
type AlibabaScbpTargetAdCampaignProductEffectAPIResponse struct {
    model.CommonResponse
    AlibabaScbpTargetAdCampaignProductEffectAPIResponseModel
}

// 定向推广-获取计划中产品推广效果 成功返回结果
type AlibabaScbpTargetAdCampaignProductEffectAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_scbp_target_ad_campaign_product_effect_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 产品数据
    ProductList   []TopP4pQuickProductEffectView `json:"product_list,omitempty" xml:"product_list>top_p4p_quick_product_effect_view,omitempty"`
    // 总页数
    TotalPage   int64 `json:"total_page,omitempty" xml:"total_page,omitempty"`
    // 总个数
    TotalNum   int64 `json:"total_num,omitempty" xml:"total_num,omitempty"`
}
