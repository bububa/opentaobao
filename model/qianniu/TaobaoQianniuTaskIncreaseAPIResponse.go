package qianniu

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQianniuTaskIncreaseAPIResponse 增加任务接收人接口 API返回值
// taobao.qianniu.task.increase
//
// 根据任务元id增加任务接收人
type TaobaoQianniuTaskIncreaseAPIResponse struct {
	model.CommonResponse
	TaobaoQianniuTaskIncreaseAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoQianniuTaskIncreaseAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoQianniuTaskIncreaseAPIResponseModel).Reset()
}

// TaobaoQianniuTaskIncreaseAPIResponseModel is 增加任务接收人接口 成功返回结果
type TaobaoQianniuTaskIncreaseAPIResponseModel struct {
	XMLName xml.Name `xml:"qianniu_task_increase_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 是否添加成功
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoQianniuTaskIncreaseAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = false
}

var poolTaobaoQianniuTaskIncreaseAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoQianniuTaskIncreaseAPIResponse)
	},
}

// GetTaobaoQianniuTaskIncreaseAPIResponse 从 sync.Pool 获取 TaobaoQianniuTaskIncreaseAPIResponse
func GetTaobaoQianniuTaskIncreaseAPIResponse() *TaobaoQianniuTaskIncreaseAPIResponse {
	return poolTaobaoQianniuTaskIncreaseAPIResponse.Get().(*TaobaoQianniuTaskIncreaseAPIResponse)
}

// ReleaseTaobaoQianniuTaskIncreaseAPIResponse 将 TaobaoQianniuTaskIncreaseAPIResponse 保存到 sync.Pool
func ReleaseTaobaoQianniuTaskIncreaseAPIResponse(v *TaobaoQianniuTaskIncreaseAPIResponse) {
	v.Reset()
	poolTaobaoQianniuTaskIncreaseAPIResponse.Put(v)
}
