package simba

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaCampaignScheduleGetAPIResponse 取得一个推广计划的分时折扣设置 API返回值
// taobao.simba.campaign.schedule.get
//
// 取得一个推广计划的分时折扣设置
type TaobaoSimbaCampaignScheduleGetAPIResponse struct {
	model.CommonResponse
	TaobaoSimbaCampaignScheduleGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoSimbaCampaignScheduleGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoSimbaCampaignScheduleGetAPIResponseModel).Reset()
}

// TaobaoSimbaCampaignScheduleGetAPIResponseModel is 取得一个推广计划的分时折扣设置 成功返回结果
type TaobaoSimbaCampaignScheduleGetAPIResponseModel struct {
	XMLName xml.Name `xml:"simba_campaign_schedule_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 修改后的推广计划分时折扣
	CampaignSchedule *CampaignSchedule `json:"campaign_schedule,omitempty" xml:"campaign_schedule,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoSimbaCampaignScheduleGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.CampaignSchedule = nil
}

var poolTaobaoSimbaCampaignScheduleGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoSimbaCampaignScheduleGetAPIResponse)
	},
}

// GetTaobaoSimbaCampaignScheduleGetAPIResponse 从 sync.Pool 获取 TaobaoSimbaCampaignScheduleGetAPIResponse
func GetTaobaoSimbaCampaignScheduleGetAPIResponse() *TaobaoSimbaCampaignScheduleGetAPIResponse {
	return poolTaobaoSimbaCampaignScheduleGetAPIResponse.Get().(*TaobaoSimbaCampaignScheduleGetAPIResponse)
}

// ReleaseTaobaoSimbaCampaignScheduleGetAPIResponse 将 TaobaoSimbaCampaignScheduleGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoSimbaCampaignScheduleGetAPIResponse(v *TaobaoSimbaCampaignScheduleGetAPIResponse) {
	v.Reset()
	poolTaobaoSimbaCampaignScheduleGetAPIResponse.Put(v)
}
