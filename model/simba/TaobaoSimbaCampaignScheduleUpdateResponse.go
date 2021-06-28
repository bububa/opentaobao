package simba

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
更新一个推广计划的分时折扣设置 APIResponse
taobao.simba.campaign.schedule.update

更新一个推广计划的分时折扣设置
*/
type TaobaoSimbaCampaignScheduleUpdateAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"simba_campaign_schedule_update_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 修改后的推广计划分时折扣
    
    CampaignSchedule   *CampaignSchedule `json:"campaign_schedule,omitempty" xml:"