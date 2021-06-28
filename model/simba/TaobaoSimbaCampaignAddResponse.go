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
	RequestId     string         `json:"request_id,omitempty" xml:"simba_campaign_add_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 创建的推广计划
    
    Campaign   *Campaign `json:"campaign,omitempty" xml:"