package qianniu

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQianniuTaskCancelAPIResponse 取消轻任务 API返回值
// taobao.qianniu.task.cancel
//
// 由任务发起者调用
type TaobaoQianniuTaskCancelAPIResponse struct {
	model.CommonResponse
	TaobaoQianniuTaskCancelAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoQianniuTaskCancelAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoQianniuTaskCancelAPIResponseModel).Reset()
}

// TaobaoQianniuTaskCancelAPIResponseModel is 取消轻任务 成功返回结果
type TaobaoQianniuTaskCancelAPIResponseModel struct {
	XMLName xml.Name `xml:"qianniu_task_cancel_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 是否成功
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoQianniuTaskCancelAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = false
}

var poolTaobaoQianniuTaskCancelAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoQianniuTaskCancelAPIResponse)
	},
}

// GetTaobaoQianniuTaskCancelAPIResponse 从 sync.Pool 获取 TaobaoQianniuTaskCancelAPIResponse
func GetTaobaoQianniuTaskCancelAPIResponse() *TaobaoQianniuTaskCancelAPIResponse {
	return poolTaobaoQianniuTaskCancelAPIResponse.Get().(*TaobaoQianniuTaskCancelAPIResponse)
}

// ReleaseTaobaoQianniuTaskCancelAPIResponse 将 TaobaoQianniuTaskCancelAPIResponse 保存到 sync.Pool
func ReleaseTaobaoQianniuTaskCancelAPIResponse(v *TaobaoQianniuTaskCancelAPIResponse) {
	v.Reset()
	poolTaobaoQianniuTaskCancelAPIResponse.Put(v)
}
