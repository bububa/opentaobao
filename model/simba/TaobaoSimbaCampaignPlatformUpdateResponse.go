package simba

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
更新一个推广计划的平台设置 APIResponse
taobao.simba.campaign.platform.update

更新一个推广计划的平台设置
*/
type TaobaoSimbaCampaignPlatformUpdateAPIResponse struct {
    model.CommonResponse
    TaobaoSimbaCampaignPlatformUpdateResponse
}

type TaobaoSimbaCampaignPlatformUpdateResponse struct {
    XMLName xml.Name `xml:"simba_campaign_platform_update_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 修改后的推广计划投放平台设置
    
    CampaignPlatform   *CampaignPlatform `json:"campaign_platform,omitempty" xml:"campaign_platform,omitempty"`

    
}
