package servicecenter

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWeikeEserviceScheduleGetAPIResponse 客服排班信息查询接口 API返回值
// taobao.weike.eservice.schedule.get
//
// 客服排班信息查询接口
type TaobaoWeikeEserviceScheduleGetAPIResponse struct {
	model.CommonResponse
	TaobaoWeikeEserviceScheduleGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoWeikeEserviceScheduleGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoWeikeEserviceScheduleGetAPIResponseModel).Reset()
}

// TaobaoWeikeEserviceScheduleGetAPIResponseModel is 客服排班信息查询接口 成功返回结果
type TaobaoWeikeEserviceScheduleGetAPIResponseModel struct {
	XMLName xml.Name `xml:"weike_eservice_schedule_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 排班信息查询结果
	Result *CsSchedulingWrapper `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoWeikeEserviceScheduleGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoWeikeEserviceScheduleGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoWeikeEserviceScheduleGetAPIResponse)
	},
}

// GetTaobaoWeikeEserviceScheduleGetAPIResponse 从 sync.Pool 获取 TaobaoWeikeEserviceScheduleGetAPIResponse
func GetTaobaoWeikeEserviceScheduleGetAPIResponse() *TaobaoWeikeEserviceScheduleGetAPIResponse {
	return poolTaobaoWeikeEserviceScheduleGetAPIResponse.Get().(*TaobaoWeikeEserviceScheduleGetAPIResponse)
}

// ReleaseTaobaoWeikeEserviceScheduleGetAPIResponse 将 TaobaoWeikeEserviceScheduleGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoWeikeEserviceScheduleGetAPIResponse(v *TaobaoWeikeEserviceScheduleGetAPIResponse) {
	v.Reset()
	poolTaobaoWeikeEserviceScheduleGetAPIResponse.Put(v)
}
