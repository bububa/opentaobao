package scs

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOnebpDkxCampaignCampaignNoreportAPIResponse 获取场景计划的非报表数据 API返回值
// taobao.onebp.dkx.campaign.campaign.noreport
//
// 获取场景计划的非报表数据。场景和bizCode的对应关系为：拉新快adStrategyDkx，上新快adStrategyShangXin ，货品加速adStrategyProductSpeed，入会快adStrategyRuHui，预热蓄水adStrategyYuRe，爆发收割adStrategyBaoFa
type TaobaoOnebpDkxCampaignCampaignNoreportAPIResponse struct {
	model.CommonResponse
	TaobaoOnebpDkxCampaignCampaignNoreportAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoOnebpDkxCampaignCampaignNoreportAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoOnebpDkxCampaignCampaignNoreportAPIResponseModel).Reset()
}

// TaobaoOnebpDkxCampaignCampaignNoreportAPIResponseModel is 获取场景计划的非报表数据 成功返回结果
type TaobaoOnebpDkxCampaignCampaignNoreportAPIResponseModel struct {
	XMLName xml.Name `xml:"onebp_dkx_campaign_campaign_noreport_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果体
	Result *TaobaoOnebpDkxCampaignCampaignNoreportResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoOnebpDkxCampaignCampaignNoreportAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoOnebpDkxCampaignCampaignNoreportAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoOnebpDkxCampaignCampaignNoreportAPIResponse)
	},
}

// GetTaobaoOnebpDkxCampaignCampaignNoreportAPIResponse 从 sync.Pool 获取 TaobaoOnebpDkxCampaignCampaignNoreportAPIResponse
func GetTaobaoOnebpDkxCampaignCampaignNoreportAPIResponse() *TaobaoOnebpDkxCampaignCampaignNoreportAPIResponse {
	return poolTaobaoOnebpDkxCampaignCampaignNoreportAPIResponse.Get().(*TaobaoOnebpDkxCampaignCampaignNoreportAPIResponse)
}

// ReleaseTaobaoOnebpDkxCampaignCampaignNoreportAPIResponse 将 TaobaoOnebpDkxCampaignCampaignNoreportAPIResponse 保存到 sync.Pool
func ReleaseTaobaoOnebpDkxCampaignCampaignNoreportAPIResponse(v *TaobaoOnebpDkxCampaignCampaignNoreportAPIResponse) {
	v.Reset()
	poolTaobaoOnebpDkxCampaignCampaignNoreportAPIResponse.Put(v)
}
