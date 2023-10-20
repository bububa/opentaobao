package qianniu

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQianniuTaskCreateAPIResponse 创建轻任务 API返回值
// taobao.qianniu.task.create
//
// 发起一个轻任务，分配给多个执行者，并发送消息提醒，由任务发起者调用
type TaobaoQianniuTaskCreateAPIResponse struct {
	model.CommonResponse
	TaobaoQianniuTaskCreateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoQianniuTaskCreateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoQianniuTaskCreateAPIResponseModel).Reset()
}

// TaobaoQianniuTaskCreateAPIResponseModel is 创建轻任务 成功返回结果
type TaobaoQianniuTaskCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"qianniu_task_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 创建的任务元数据
	Result *QTaskMetadata `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoQianniuTaskCreateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoQianniuTaskCreateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoQianniuTaskCreateAPIResponse)
	},
}

// GetTaobaoQianniuTaskCreateAPIResponse 从 sync.Pool 获取 TaobaoQianniuTaskCreateAPIResponse
func GetTaobaoQianniuTaskCreateAPIResponse() *TaobaoQianniuTaskCreateAPIResponse {
	return poolTaobaoQianniuTaskCreateAPIResponse.Get().(*TaobaoQianniuTaskCreateAPIResponse)
}

// ReleaseTaobaoQianniuTaskCreateAPIResponse 将 TaobaoQianniuTaskCreateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoQianniuTaskCreateAPIResponse(v *TaobaoQianniuTaskCreateAPIResponse) {
	v.Reset()
	poolTaobaoQianniuTaskCreateAPIResponse.Put(v)
}
