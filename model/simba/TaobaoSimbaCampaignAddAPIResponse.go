package simba

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaCampaignAddAPIResponse 创建一个推广计划 API返回值
// taobao.simba.campaign.add
//
// 创建一个推广计划
type TaobaoSimbaCampaignAddAPIResponse struct {
	model.CommonResponse
	TaobaoSimbaCampaignAddAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoSimbaCampaignAddAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoSimbaCampaignAddAPIResponseModel).Reset()
}

// TaobaoSimbaCampaignAddAPIResponseModel is 创建一个推广计划 成功返回结果
type TaobaoSimbaCampaignAddAPIResponseModel struct {
	XMLName xml.Name `xml:"simba_campaign_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 创建的推广计划
	Campaign *Campaign `json:"campaign,omitempty" xml:"campaign,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoSimbaCampaignAddAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Campaign = nil
}

var poolTaobaoSimbaCampaignAddAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoSimbaCampaignAddAPIResponse)
	},
}

// GetTaobaoSimbaCampaignAddAPIResponse 从 sync.Pool 获取 TaobaoSimbaCampaignAddAPIResponse
func GetTaobaoSimbaCampaignAddAPIResponse() *TaobaoSimbaCampaignAddAPIResponse {
	return poolTaobaoSimbaCampaignAddAPIResponse.Get().(*TaobaoSimbaCampaignAddAPIResponse)
}

// ReleaseTaobaoSimbaCampaignAddAPIResponse 将 TaobaoSimbaCampaignAddAPIResponse 保存到 sync.Pool
func ReleaseTaobaoSimbaCampaignAddAPIResponse(v *TaobaoSimbaCampaignAddAPIResponse) {
	v.Reset()
	poolTaobaoSimbaCampaignAddAPIResponse.Put(v)
}
