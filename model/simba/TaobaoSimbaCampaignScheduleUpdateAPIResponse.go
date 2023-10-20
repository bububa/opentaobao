package simba

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaCampaignScheduleUpdateAPIResponse 更新一个推广计划的分时折扣设置 API返回值
// taobao.simba.campaign.schedule.update
//
// 更新一个推广计划的分时折扣设置
type TaobaoSimbaCampaignScheduleUpdateAPIResponse struct {
	model.CommonResponse
	TaobaoSimbaCampaignScheduleUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoSimbaCampaignScheduleUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoSimbaCampaignScheduleUpdateAPIResponseModel).Reset()
}

// TaobaoSimbaCampaignScheduleUpdateAPIResponseModel is 更新一个推广计划的分时折扣设置 成功返回结果
type TaobaoSimbaCampaignScheduleUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"simba_campaign_schedule_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 修改后的推广计划分时折扣
	CampaignSchedule *CampaignSchedule `json:"campaign_schedule,omitempty" xml:"campaign_schedule,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoSimbaCampaignScheduleUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.CampaignSchedule = nil
}

var poolTaobaoSimbaCampaignScheduleUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoSimbaCampaignScheduleUpdateAPIResponse)
	},
}

// GetTaobaoSimbaCampaignScheduleUpdateAPIResponse 从 sync.Pool 获取 TaobaoSimbaCampaignScheduleUpdateAPIResponse
func GetTaobaoSimbaCampaignScheduleUpdateAPIResponse() *TaobaoSimbaCampaignScheduleUpdateAPIResponse {
	return poolTaobaoSimbaCampaignScheduleUpdateAPIResponse.Get().(*TaobaoSimbaCampaignScheduleUpdateAPIResponse)
}

// ReleaseTaobaoSimbaCampaignScheduleUpdateAPIResponse 将 TaobaoSimbaCampaignScheduleUpdateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoSimbaCampaignScheduleUpdateAPIResponse(v *TaobaoSimbaCampaignScheduleUpdateAPIResponse) {
	v.Reset()
	poolTaobaoSimbaCampaignScheduleUpdateAPIResponse.Put(v)
}
