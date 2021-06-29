package simba

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
更新一个推广计划 API返回值 
taobao.simba.campaign.update

更新一个推广计划，可以设置推广计划名字，修改推广计划上下线状态。
*/
type TaobaoSimbaCampaignUpdateAPIResponse struct {
    model.CommonResponse
    TaobaoSimbaCampaignUpdateResponse
}

// 更新一个推广计划 成功返回结果
type TaobaoSimbaCampaignUpdateResponse struct {
    XMLName xml.Name `xml:"simba_campaign_update_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 修改后的推广计划
    Campaign   *Campaign `json:"campaign,omitempty" xml:"campaign,omitempty"`
}
