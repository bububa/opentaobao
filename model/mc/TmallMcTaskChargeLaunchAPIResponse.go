package mc

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallMcTaskChargeLaunchAPIResponse 云码充电宝投放链路 API返回值
// tmall.mc.task.charge.launch
//
// 云码充电宝投放链路，用于判断用户是否有匹配的投放计划
type TmallMcTaskChargeLaunchAPIResponse struct {
	model.CommonResponse
	TmallMcTaskChargeLaunchAPIResponseModel
}

// Reset 清空结构体
func (m *TmallMcTaskChargeLaunchAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallMcTaskChargeLaunchAPIResponseModel).Reset()
}

// TmallMcTaskChargeLaunchAPIResponseModel is 云码充电宝投放链路 成功返回结果
type TmallMcTaskChargeLaunchAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_mc_task_charge_launch_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 入会页url
	Result string `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallMcTaskChargeLaunchAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = ""
}

var poolTmallMcTaskChargeLaunchAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallMcTaskChargeLaunchAPIResponse)
	},
}

// GetTmallMcTaskChargeLaunchAPIResponse 从 sync.Pool 获取 TmallMcTaskChargeLaunchAPIResponse
func GetTmallMcTaskChargeLaunchAPIResponse() *TmallMcTaskChargeLaunchAPIResponse {
	return poolTmallMcTaskChargeLaunchAPIResponse.Get().(*TmallMcTaskChargeLaunchAPIResponse)
}

// ReleaseTmallMcTaskChargeLaunchAPIResponse 将 TmallMcTaskChargeLaunchAPIResponse 保存到 sync.Pool
func ReleaseTmallMcTaskChargeLaunchAPIResponse(v *TmallMcTaskChargeLaunchAPIResponse) {
	v.Reset()
	poolTmallMcTaskChargeLaunchAPIResponse.Put(v)
}
