package qianniu

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQianniuTaskUpdateAPIResponse 更新轻任务 API返回值
// taobao.qianniu.task.update
//
// 由任务执行者调用，sub_status，tag和memo至少提供一个
type TaobaoQianniuTaskUpdateAPIResponse struct {
	model.CommonResponse
	TaobaoQianniuTaskUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoQianniuTaskUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoQianniuTaskUpdateAPIResponseModel).Reset()
}

// TaobaoQianniuTaskUpdateAPIResponseModel is 更新轻任务 成功返回结果
type TaobaoQianniuTaskUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"qianniu_task_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 是否成功
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoQianniuTaskUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = false
}

var poolTaobaoQianniuTaskUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoQianniuTaskUpdateAPIResponse)
	},
}

// GetTaobaoQianniuTaskUpdateAPIResponse 从 sync.Pool 获取 TaobaoQianniuTaskUpdateAPIResponse
func GetTaobaoQianniuTaskUpdateAPIResponse() *TaobaoQianniuTaskUpdateAPIResponse {
	return poolTaobaoQianniuTaskUpdateAPIResponse.Get().(*TaobaoQianniuTaskUpdateAPIResponse)
}

// ReleaseTaobaoQianniuTaskUpdateAPIResponse 将 TaobaoQianniuTaskUpdateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoQianniuTaskUpdateAPIResponse(v *TaobaoQianniuTaskUpdateAPIResponse) {
	v.Reset()
	poolTaobaoQianniuTaskUpdateAPIResponse.Put(v)
}
