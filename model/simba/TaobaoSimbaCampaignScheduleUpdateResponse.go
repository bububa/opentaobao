package simba

import (
    "github.com/bububa/opentaobao/model"
)

/* 
更新一个推广计划的分时折扣设置 APIResponse
taobao.simba.campaign.schedule.update

更新一个推广计划的分时折扣设置
*/
type TaobaoSimbaCampaignScheduleUpdateAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoSimbaCampaignScheduleUpdateResponse `json:"simba_campaign_schedule_update_response,omitempty"` 
    TaobaoSimbaCampaignScheduleUpdateResponse
}

/* model for simplify = false
type TaobaoSimbaCampaignScheduleUpdateResponse struct {

    // 修改后的推广计划分时折扣
    
    CampaignSchedule  *struct {
        CampaignSchedule  *CampaignSchedule `json:"campaign_schedule,omitempty"`
    } `json:"campaign_schedule,omitempty"`
    

}
*/

type TaobaoSimbaCampaignScheduleUpdateResponse struct {

    // 修改后的推广计划分时折扣
    CampaignSchedule   *CampaignSchedule `json:"campaign_schedule,omitempty"`

}
