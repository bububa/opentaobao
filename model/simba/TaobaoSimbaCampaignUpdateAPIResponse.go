package simba

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaCampaignUpdateAPIResponse 更新一个推广计划 API返回值
// taobao.simba.campaign.update
//
// 更新一个推广计划，可以设置推广计划名字，修改推广计划上下线状态。
type TaobaoSimbaCampaignUpdateAPIResponse struct {
	model.CommonResponse
	TaobaoSimbaCampaignUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoSimbaCampaignUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoSimbaCampaignUpdateAPIResponseModel).Reset()
}

// TaobaoSimbaCampaignUpdateAPIResponseModel is 更新一个推广计划 成功返回结果
type TaobaoSimbaCampaignUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"simba_campaign_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 修改后的推广计划
	Campaign *Campaign `json:"campaign,omitempty" xml:"campaign,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoSimbaCampaignUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Campaign = nil
}

var poolTaobaoSimbaCampaignUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoSimbaCampaignUpdateAPIResponse)
	},
}

// GetTaobaoSimbaCampaignUpdateAPIResponse 从 sync.Pool 获取 TaobaoSimbaCampaignUpdateAPIResponse
func GetTaobaoSimbaCampaignUpdateAPIResponse() *TaobaoSimbaCampaignUpdateAPIResponse {
	return poolTaobaoSimbaCampaignUpdateAPIResponse.Get().(*TaobaoSimbaCampaignUpdateAPIResponse)
}

// ReleaseTaobaoSimbaCampaignUpdateAPIResponse 将 TaobaoSimbaCampaignUpdateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoSimbaCampaignUpdateAPIResponse(v *TaobaoSimbaCampaignUpdateAPIResponse) {
	v.Reset()
	poolTaobaoSimbaCampaignUpdateAPIResponse.Put(v)
}
