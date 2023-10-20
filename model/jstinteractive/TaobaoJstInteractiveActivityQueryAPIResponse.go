package jstinteractive

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoJstInteractiveActivityQueryAPIResponse 互动任务活动查询接口 API返回值
// taobao.jst.interactive.activity.query
//
// 互动任务活动查询接口
type TaobaoJstInteractiveActivityQueryAPIResponse struct {
	model.CommonResponse
	TaobaoJstInteractiveActivityQueryAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoJstInteractiveActivityQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoJstInteractiveActivityQueryAPIResponseModel).Reset()
}

// TaobaoJstInteractiveActivityQueryAPIResponseModel is 互动任务活动查询接口 成功返回结果
type TaobaoJstInteractiveActivityQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"jst_interactive_activity_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 活动列表，只返回有效的活动
	ActivityList []Activity `json:"activity_list,omitempty" xml:"activity_list>activity,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoJstInteractiveActivityQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ActivityList = m.ActivityList[:0]
}

var poolTaobaoJstInteractiveActivityQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoJstInteractiveActivityQueryAPIResponse)
	},
}

// GetTaobaoJstInteractiveActivityQueryAPIResponse 从 sync.Pool 获取 TaobaoJstInteractiveActivityQueryAPIResponse
func GetTaobaoJstInteractiveActivityQueryAPIResponse() *TaobaoJstInteractiveActivityQueryAPIResponse {
	return poolTaobaoJstInteractiveActivityQueryAPIResponse.Get().(*TaobaoJstInteractiveActivityQueryAPIResponse)
}

// ReleaseTaobaoJstInteractiveActivityQueryAPIResponse 将 TaobaoJstInteractiveActivityQueryAPIResponse 保存到 sync.Pool
func ReleaseTaobaoJstInteractiveActivityQueryAPIResponse(v *TaobaoJstInteractiveActivityQueryAPIResponse) {
	v.Reset()
	poolTaobaoJstInteractiveActivityQueryAPIResponse.Put(v)
}
