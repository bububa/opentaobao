package simba

import (
    "github.com/bububa/opentaobao/model"
)

/* 
取得一个推广计划的分时折扣设置 APIResponse
taobao.simba.campaign.schedule.get

取得一个推广计划的分时折扣设置
*/
type TaobaoSimbaCampaignScheduleGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoSimbaCampaignScheduleGetResponse `json:"simba_campaign_schedule_get_response,omitempty"` 
    TaobaoSimbaCampaignScheduleGetResponse
}

/* model for simplify = false
type TaobaoSimbaCampaignScheduleGetResponse struct {

    // 修改后的推广计划分时折扣
    
    CampaignSchedule  *struct {
        CampaignSchedule  *CampaignSchedule `json:"campaign_schedule,omitempty"`
    } `json:"campaign_schedule,omitempty"`
    

}
*/

type TaobaoSimbaCampaignScheduleGetResponse struct {

    // 修改后的推广计划分时折扣
    CampaignSchedule   *CampaignSchedule `json:"campaign_schedule,omitempty"`

}
