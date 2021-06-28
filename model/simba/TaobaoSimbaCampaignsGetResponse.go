package simba

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
取得一组推广计划 APIResponse
taobao.simba.campaigns.get

取得一个客户的推广计划；
*/
type TaobaoSimbaCampaignsGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"simba_campaigns_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 推广计划列表
    
    Campaigns   []Campaign `json:"campaigns,omitempty" xml:"