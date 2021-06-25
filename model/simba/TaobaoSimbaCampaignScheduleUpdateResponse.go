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
    Response *TaobaoSimbaCampaignScheduleUpdateResponse `json:"taobao_simba_campaign_schedule_update_response,omitempty"`
}

type TaobaoSimbaCampaignScheduleUpdateResponse struct {

    // 修改后的推广计划分时折扣
    CampaignSchedule   *CampaignSchedule `json:"campaign_schedule,omitempty"`

}
