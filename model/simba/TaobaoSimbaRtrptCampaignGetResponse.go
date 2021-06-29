package simba

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取推广计划实时报表数据 APIResponse
taobao.simba.rtrpt.campaign.get

获取推广计划实时报表数据
*/
type TaobaoSimbaRtrptCampaignGetAPIResponse struct {
    model.CommonResponse
    TaobaoSimbaRtrptCampaignGetResponse
}

type TaobaoSimbaRtrptCampaignGetResponse struct {
    XMLName xml.Name `xml:"simba_rtrpt_campaign_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 111
    
    Resultss   []RtRptResultEntityDTO `json:"resultss,omitempty" xml:"resultss>rt_rpt_result_entity_dto,omitempty"`
    
    
}
