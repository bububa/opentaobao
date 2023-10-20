package scs

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoonebpdkxcampaigncampaignmodifyAPIResponse 修改计划 API返回值
// taobao.onebp.dkx.campaign.campaign.modify
//
// 修改计划。场景和bizCode的对应关系为：拉新快adStrategyDkx，上新快adStrategyShangXin ，货品加速adStrategyProductSpeed，入会快adStrategyRuHui，预热蓄水adStrategyYuRe，爆发收割adStrategyBaoFa
type TaobaoonebpdkxcampaigncampaignmodifyAPIResponse struct {
	model.CommonResponse
	TaobaoonebpdkxcampaigncampaignmodifyAPIResponseModel
}

// TaobaoonebpdkxcampaigncampaignmodifyAPIResponseModel is 修改计划 成功返回结果
type TaobaoonebpdkxcampaigncampaignmodifyAPIResponseModel struct {
	XMLName xml.Name `xml:"onebp_dkx_campaign_campaign_modify_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果体
	Result *TaobaoonebpdkxcampaigncampaignmodifyResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
