package qianniu

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQianniuTaskRemoveAPIResponse 轻任务删除接口 API返回值
// taobao.qianniu.task.remove
//
// 轻任务删除接口。
type TaobaoQianniuTaskRemoveAPIResponse struct {
	model.CommonResponse
	TaobaoQianniuTaskRemoveAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoQianniuTaskRemoveAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoQianniuTaskRemoveAPIResponseModel).Reset()
}

// TaobaoQianniuTaskRemoveAPIResponseModel is 轻任务删除接口 成功返回结果
type TaobaoQianniuTaskRemoveAPIResponseModel struct {
	XMLName xml.Name `xml:"qianniu_task_remove_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 是否成功
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoQianniuTaskRemoveAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = false
}

var poolTaobaoQianniuTaskRemoveAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoQianniuTaskRemoveAPIResponse)
	},
}

// GetTaobaoQianniuTaskRemoveAPIResponse 从 sync.Pool 获取 TaobaoQianniuTaskRemoveAPIResponse
func GetTaobaoQianniuTaskRemoveAPIResponse() *TaobaoQianniuTaskRemoveAPIResponse {
	return poolTaobaoQianniuTaskRemoveAPIResponse.Get().(*TaobaoQianniuTaskRemoveAPIResponse)
}

// ReleaseTaobaoQianniuTaskRemoveAPIResponse 将 TaobaoQianniuTaskRemoveAPIResponse 保存到 sync.Pool
func ReleaseTaobaoQianniuTaskRemoveAPIResponse(v *TaobaoQianniuTaskRemoveAPIResponse) {
	v.Reset()
	poolTaobaoQianniuTaskRemoveAPIResponse.Put(v)
}
