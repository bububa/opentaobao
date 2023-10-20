package eticket

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoVmarketEticketTasksGetAPIResponse 任务列表获取接口 API返回值
// taobao.vmarket.eticket.tasks.get
//
// 外部合作卖家获取任务列表的信息：如发码同通知失败或者回调失败的订单号
type TaobaoVmarketEticketTasksGetAPIResponse struct {
	model.CommonResponse
	TaobaoVmarketEticketTasksGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoVmarketEticketTasksGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoVmarketEticketTasksGetAPIResponseModel).Reset()
}

// TaobaoVmarketEticketTasksGetAPIResponseModel is 任务列表获取接口 成功返回结果
type TaobaoVmarketEticketTasksGetAPIResponseModel struct {
	XMLName xml.Name `xml:"vmarket_eticket_tasks_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 任务列表查询结果信息
	EticketTasks []EticketTask `json:"eticket_tasks,omitempty" xml:"eticket_tasks>eticket_task,omitempty"`
	// 任务列表查询结果的总数
	TotalResults int64 `json:"total_results,omitempty" xml:"total_results,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoVmarketEticketTasksGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.EticketTasks = m.EticketTasks[:0]
	m.TotalResults = 0
}

var poolTaobaoVmarketEticketTasksGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoVmarketEticketTasksGetAPIResponse)
	},
}

// GetTaobaoVmarketEticketTasksGetAPIResponse 从 sync.Pool 获取 TaobaoVmarketEticketTasksGetAPIResponse
func GetTaobaoVmarketEticketTasksGetAPIResponse() *TaobaoVmarketEticketTasksGetAPIResponse {
	return poolTaobaoVmarketEticketTasksGetAPIResponse.Get().(*TaobaoVmarketEticketTasksGetAPIResponse)
}

// ReleaseTaobaoVmarketEticketTasksGetAPIResponse 将 TaobaoVmarketEticketTasksGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoVmarketEticketTasksGetAPIResponse(v *TaobaoVmarketEticketTasksGetAPIResponse) {
	v.Reset()
	poolTaobaoVmarketEticketTasksGetAPIResponse.Put(v)
}
