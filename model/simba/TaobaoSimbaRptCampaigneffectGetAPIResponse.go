package simba

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaosimbarptcampaigneffectgetAPIResponse 推广计划效果报表数据对象 API返回值
// taobao.simba.rpt.campaigneffect.get
//
// 推广计划效果报表数据对象
type TaobaosimbarptcampaigneffectgetAPIResponse struct {
	model.CommonResponse
	TaobaosimbarptcampaigneffectgetAPIResponseModel
}

// TaobaosimbarptcampaigneffectgetAPIResponseModel is 推广计划效果报表数据对象 成功返回结果
type TaobaosimbarptcampaigneffectgetAPIResponseModel struct {
	XMLName xml.Name `xml:"simba_rpt_campaigneffect_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 推广计划效果报表数据对象
	RptCampaignEffectList string `json:"rpt_campaign_effect_list,omitempty" xml:"rpt_campaign_effect_list,omitempty"`
}
