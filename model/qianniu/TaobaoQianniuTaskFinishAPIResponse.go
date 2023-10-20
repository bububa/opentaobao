package qianniu

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQianniuTaskFinishAPIResponse 完成轻任务 API返回值
// taobao.qianniu.task.finish
//
// 由任务执行者调用
type TaobaoQianniuTaskFinishAPIResponse struct {
	model.CommonResponse
	TaobaoQianniuTaskFinishAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoQianniuTaskFinishAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoQianniuTaskFinishAPIResponseModel).Reset()
}

// TaobaoQianniuTaskFinishAPIResponseModel is 完成轻任务 成功返回结果
type TaobaoQianniuTaskFinishAPIResponseModel struct {
	XMLName xml.Name `xml:"qianniu_task_finish_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 是否成功
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoQianniuTaskFinishAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = false
}

var poolTaobaoQianniuTaskFinishAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoQianniuTaskFinishAPIResponse)
	},
}

// GetTaobaoQianniuTaskFinishAPIResponse 从 sync.Pool 获取 TaobaoQianniuTaskFinishAPIResponse
func GetTaobaoQianniuTaskFinishAPIResponse() *TaobaoQianniuTaskFinishAPIResponse {
	return poolTaobaoQianniuTaskFinishAPIResponse.Get().(*TaobaoQianniuTaskFinishAPIResponse)
}

// ReleaseTaobaoQianniuTaskFinishAPIResponse 将 TaobaoQianniuTaskFinishAPIResponse 保存到 sync.Pool
func ReleaseTaobaoQianniuTaskFinishAPIResponse(v *TaobaoQianniuTaskFinishAPIResponse) {
	v.Reset()
	poolTaobaoQianniuTaskFinishAPIResponse.Put(v)
}
