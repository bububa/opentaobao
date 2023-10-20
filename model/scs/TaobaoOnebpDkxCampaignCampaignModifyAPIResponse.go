package scs

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOnebpDkxCampaignCampaignModifyAPIResponse 修改计划 API返回值
// taobao.onebp.dkx.campaign.campaign.modify
//
// 修改计划。场景和bizCode的对应关系为：拉新快adStrategyDkx，上新快adStrategyShangXin ，货品加速adStrategyProductSpeed，入会快adStrategyRuHui，预热蓄水adStrategyYuRe，爆发收割adStrategyBaoFa
type TaobaoOnebpDkxCampaignCampaignModifyAPIResponse struct {
	model.CommonResponse
	TaobaoOnebpDkxCampaignCampaignModifyAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoOnebpDkxCampaignCampaignModifyAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoOnebpDkxCampaignCampaignModifyAPIResponseModel).Reset()
}

// TaobaoOnebpDkxCampaignCampaignModifyAPIResponseModel is 修改计划 成功返回结果
type TaobaoOnebpDkxCampaignCampaignModifyAPIResponseModel struct {
	XMLName xml.Name `xml:"onebp_dkx_campaign_campaign_modify_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果体
	Result *TaobaoOnebpDkxCampaignCampaignModifyResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoOnebpDkxCampaignCampaignModifyAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoOnebpDkxCampaignCampaignModifyAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoOnebpDkxCampaignCampaignModifyAPIResponse)
	},
}

// GetTaobaoOnebpDkxCampaignCampaignModifyAPIResponse 从 sync.Pool 获取 TaobaoOnebpDkxCampaignCampaignModifyAPIResponse
func GetTaobaoOnebpDkxCampaignCampaignModifyAPIResponse() *TaobaoOnebpDkxCampaignCampaignModifyAPIResponse {
	return poolTaobaoOnebpDkxCampaignCampaignModifyAPIResponse.Get().(*TaobaoOnebpDkxCampaignCampaignModifyAPIResponse)
}

// ReleaseTaobaoOnebpDkxCampaignCampaignModifyAPIResponse 将 TaobaoOnebpDkxCampaignCampaignModifyAPIResponse 保存到 sync.Pool
func ReleaseTaobaoOnebpDkxCampaignCampaignModifyAPIResponse(v *TaobaoOnebpDkxCampaignCampaignModifyAPIResponse) {
	v.Reset()
	poolTaobaoOnebpDkxCampaignCampaignModifyAPIResponse.Put(v)
}
