package simba

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
更新一个推广计划的投放地域 API返回值 
taobao.simba.campaign.area.update

更新一个推广计划的投放地域
*/
type TaobaoSimbaCampaignAreaUpdateAPIResponse struct {
    model.CommonResponse
    TaobaoSimbaCampaignAreaUpdateAPIResponseModel
}

// 更新一个推广计划的投放地域 成功返回结果
type TaobaoSimbaCampaignAreaUpdateAPIResponseModel struct {
    XMLName xml.Name `xml:"simba_campaign_area_update_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 修改后的推广计划投放地域
    CampaignArea   *CampaignArea `json:"campaign_area,omitempty" xml:"campaign_area,omitempty"`
}
