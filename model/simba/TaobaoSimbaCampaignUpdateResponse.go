package simba

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
更新一个推广计划 APIResponse
taobao.simba.campaign.update

更新一个推广计划，可以设置推广计划名字，修改推广计划上下线状态。
*/
type TaobaoSimbaCampaignUpdateAPIResponse struct {
    model.CommonResponse
    TaobaoSimbaCampaignUpdateResponse
}

type TaobaoSimbaCampaignUpdateResponse struct {
    XMLName xml.Name `xml:"simba_campaign_update_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 修改后的推广计划
    
    Campaign   *Campaign `json:"campaign,omitempty" xml:"campaign,omitempty"`

    
}
