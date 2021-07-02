package simba

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaCampaignPlatformUpdateAPIResponse 更新一个推广计划的平台设置 API返回值
// taobao.simba.campaign.platform.update
//
// 更新一个推广计划的平台设置
type TaobaoSimbaCampaignPlatformUpdateAPIResponse struct {
	model.CommonResponse
	TaobaoSimbaCampaignPlatformUpdateAPIResponseModel
}

// TaobaoSimbaCampaignPlatformUpdateAPIResponseModel is 更新一个推广计划的平台设置 成功返回结果
type TaobaoSimbaCampaignPlatformUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"simba_campaign_platform_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 修改后的推广计划投放平台设置
	CampaignPlatform *CampaignPlatform `json:"campaign_platform,omitempty" xml:"campaign_platform,omitempty"`
}
