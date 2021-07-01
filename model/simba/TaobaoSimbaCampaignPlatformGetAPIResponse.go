package simba

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
取得一个推广计划的投放平台设置 API返回值 
taobao.simba.campaign.platform.get

获得一个推广计划的投放平台设置
*/
type TaobaoSimbaCampaignPlatformGetAPIResponse struct {
    model.CommonResponse
    TaobaoSimbaCampaignPlatformGetAPIResponseModel
}

// 取得一个推广计划的投放平台设置 成功返回结果
type TaobaoSimbaCampaignPlatformGetAPIResponseModel struct {
    XMLName xml.Name `xml:"simba_campaign_platform_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 取得的推广计划的投放平台设置
    CampaignPlatform   *CampaignPlatform `json:"campaign_platform,omitempty" xml:"campaign_platform,omitempty"`
}
