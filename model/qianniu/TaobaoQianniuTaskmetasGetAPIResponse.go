package qianniu

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQianniuTaskmetasGetAPIResponse 任务元查询接口 API返回值
// taobao.qianniu.taskmetas.get
//
// 任务元查询接口
type TaobaoQianniuTaskmetasGetAPIResponse struct {
	model.CommonResponse
	TaobaoQianniuTaskmetasGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoQianniuTaskmetasGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoQianniuTaskmetasGetAPIResponseModel).Reset()
}

// TaobaoQianniuTaskmetasGetAPIResponseModel is 任务元查询接口 成功返回结果
type TaobaoQianniuTaskmetasGetAPIResponseModel struct {
	XMLName xml.Name `xml:"qianniu_taskmetas_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// taskmetas
	Taskmetas []QTaskMetadata `json:"taskmetas,omitempty" xml:"taskmetas>q_task_metadata,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoQianniuTaskmetasGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Taskmetas = m.Taskmetas[:0]
}

var poolTaobaoQianniuTaskmetasGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoQianniuTaskmetasGetAPIResponse)
	},
}

// GetTaobaoQianniuTaskmetasGetAPIResponse 从 sync.Pool 获取 TaobaoQianniuTaskmetasGetAPIResponse
func GetTaobaoQianniuTaskmetasGetAPIResponse() *TaobaoQianniuTaskmetasGetAPIResponse {
	return poolTaobaoQianniuTaskmetasGetAPIResponse.Get().(*TaobaoQianniuTaskmetasGetAPIResponse)
}

// ReleaseTaobaoQianniuTaskmetasGetAPIResponse 将 TaobaoQianniuTaskmetasGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoQianniuTaskmetasGetAPIResponse(v *TaobaoQianniuTaskmetasGetAPIResponse) {
	v.Reset()
	poolTaobaoQianniuTaskmetasGetAPIResponse.Put(v)
}
