package simba

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
取得一个推广计划的投放地域设置 APIResponse
taobao.simba.campaign.area.get

取得一个推广计划的投放地域设置
*/
type TaobaoSimbaCampaignAreaGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"simba_campaign_area_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 推广计划的投放地域配置
    
    CampaignArea   *CampaignArea `json:"campaign_area,omitempty" xml:"