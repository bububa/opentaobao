package simba

import (
    "github.com/bububa/opentaobao/model"
)

/* 
更新一个推广计划的平台设置 APIResponse
taobao.simba.campaign.platform.update

更新一个推广计划的平台设置
*/
type TaobaoSimbaCampaignPlatformUpdateAPIResponse struct {
    model.CommonResponse
    Response *TaobaoSimbaCampaignPlatformUpdateResponse `json:"taobao_simba_campaign_platform_update_response,omitempty"`
}

type TaobaoSimbaCampaignPlatformUpdateResponse struct {

    // 修改后的推广计划投放平台设置
    CampaignPlatform   *CampaignPlatform `json:"campaign_platform,omitempty"`

}
