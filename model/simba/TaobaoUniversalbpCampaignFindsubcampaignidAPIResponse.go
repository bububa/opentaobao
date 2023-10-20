package simba

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUniversalbpCampaignFindsubcampaignidAPIResponse 查询无界版计划对应的原场景计划id API返回值
// taobao.universalbp.campaign.findsubcampaignid
//
// 查询该场景下，无界版计划对应的原场景的计划
type TaobaoUniversalbpCampaignFindsubcampaignidAPIResponse struct {
	model.CommonResponse
	TaobaoUniversalbpCampaignFindsubcampaignidAPIResponseModel
}

// TaobaoUniversalbpCampaignFindsubcampaignidAPIResponseModel is 查询无界版计划对应的原场景计划id 成功返回结果
type TaobaoUniversalbpCampaignFindsubcampaignidAPIResponseModel struct {
	XMLName xml.Name `xml:"universalbp_campaign_findsubcampaignid_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果体
	Result *TaobaoUniversalbpCampaignFindsubcampaignidTopResult `json:"result,omitempty" xml:"result,omitempty"`
}
