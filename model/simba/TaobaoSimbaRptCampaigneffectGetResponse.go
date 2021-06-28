package simba

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
推广计划效果报表数据对象 APIResponse
taobao.simba.rpt.campaigneffect.get

推广计划效果报表数据对象
*/
type TaobaoSimbaRptCampaigneffectGetAPIResponse struct {
    model.CommonResponse
    TaobaoSimbaRptCampaigneffectGetResponse
}

type TaobaoSimbaRptCampaigneffectGetResponse struct {
    XMLName xml.Name `xml:"simba_rpt_campaigneffect_get_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 推广计划效果报表数据对象
    
    RptCampaignEffectList   string `json:"rpt_campaign_effect_list,omitempty" xml:"rpt_campaign_effect_list,omitempty"`

    
}
