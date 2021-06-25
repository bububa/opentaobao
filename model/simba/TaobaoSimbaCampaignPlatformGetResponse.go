package simba

import (
    "github.com/bububa/opentaobao/model"
)

/* 
取得一个推广计划的投放平台设置 APIResponse
taobao.simba.campaign.platform.get

获得一个推广计划的投放平台设置
*/
type TaobaoSimbaCampaignPlatformGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoSimbaCampaignPlatformGetResponse `json:"taobao_simba_campaign_platform_get_response,omitempty"`
}

type TaobaoSimbaCampaignPlatformGetResponse struct {

    // 取得的推广计划的投放平台设置
    CampaignPlatform   *CampaignPlatform `json:"campaign_platform,omitempty"`

}
