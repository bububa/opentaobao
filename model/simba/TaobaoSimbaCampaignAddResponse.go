package simba

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
创建一个推广计划 APIResponse
taobao.simba.campaign.add

创建一个推广计划
*/
type TaobaoSimbaCampaignAddAPIResponse struct {
    model.CommonResponse
    TaobaoSimbaCampaignAddResponse
}

type TaobaoSimbaCampaignAddResponse struct {
    XMLName xml.Name `xml:"simba_campaign_add_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 创建的推广计划
    
    Campaign   *Campaign `json:"campaign,omitempty" xml:"campaign,omitempty"`

    
}
