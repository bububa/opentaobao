package simba

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaCampaignAreaoptionsGetAPIResponse 取得推广计划的可设置投放地域列表 API返回值
// taobao.simba.campaign.areaoptions.get
//
// 取得推广计划的可设置投放地域列表
type TaobaoSimbaCampaignAreaoptionsGetAPIResponse struct {
	model.CommonResponse
	TaobaoSimbaCampaignAreaoptionsGetAPIResponseModel
}

// TaobaoSimbaCampaignAreaoptionsGetAPIResponseModel is 取得推广计划的可设置投放地域列表 成功返回结果
type TaobaoSimbaCampaignAreaoptionsGetAPIResponseModel struct {
	XMLName xml.Name `xml:"simba_campaign_areaoptions_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 推广计划所有可设置的投放地域
	AreaOptions []AreaOption `json:"area_options,omitempty" xml:"area_options>area_option,omitempty"`
}
