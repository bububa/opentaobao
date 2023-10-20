package qianniu

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQianniuTasksCountAPIResponse 任务查询条数接口 API返回值
// taobao.qianniu.tasks.count
//
// 任务查询条数接口
type TaobaoQianniuTasksCountAPIResponse struct {
	model.CommonResponse
	TaobaoQianniuTasksCountAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoQianniuTasksCountAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoQianniuTasksCountAPIResponseModel).Reset()
}

// TaobaoQianniuTasksCountAPIResponseModel is 任务查询条数接口 成功返回结果
type TaobaoQianniuTasksCountAPIResponseModel struct {
	XMLName xml.Name `xml:"qianniu_tasks_count_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 符合查询条件的总条数
	Result int64 `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoQianniuTasksCountAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = 0
}

var poolTaobaoQianniuTasksCountAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoQianniuTasksCountAPIResponse)
	},
}

// GetTaobaoQianniuTasksCountAPIResponse 从 sync.Pool 获取 TaobaoQianniuTasksCountAPIResponse
func GetTaobaoQianniuTasksCountAPIResponse() *TaobaoQianniuTasksCountAPIResponse {
	return poolTaobaoQianniuTasksCountAPIResponse.Get().(*TaobaoQianniuTasksCountAPIResponse)
}

// ReleaseTaobaoQianniuTasksCountAPIResponse 将 TaobaoQianniuTasksCountAPIResponse 保存到 sync.Pool
func ReleaseTaobaoQianniuTasksCountAPIResponse(v *TaobaoQianniuTasksCountAPIResponse) {
	v.Reset()
	poolTaobaoQianniuTasksCountAPIResponse.Put(v)
}
