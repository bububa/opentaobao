package util

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTopEventSubscriptionQueryAPIResponse 三方事件订阅查询 API返回值
// taobao.top.event.subscription.query
//
// 三方事件订阅查询
type TaobaoTopEventSubscriptionQueryAPIResponse struct {
	model.CommonResponse
	TaobaoTopEventSubscriptionQueryAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoTopEventSubscriptionQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoTopEventSubscriptionQueryAPIResponseModel).Reset()
}

// TaobaoTopEventSubscriptionQueryAPIResponseModel is 三方事件订阅查询 成功返回结果
type TaobaoTopEventSubscriptionQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"top_event_subscription_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 是否订阅服务
	StartService bool `json:"start_service,omitempty" xml:"start_service,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoTopEventSubscriptionQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.StartService = false
}

var poolTaobaoTopEventSubscriptionQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoTopEventSubscriptionQueryAPIResponse)
	},
}

// GetTaobaoTopEventSubscriptionQueryAPIResponse 从 sync.Pool 获取 TaobaoTopEventSubscriptionQueryAPIResponse
func GetTaobaoTopEventSubscriptionQueryAPIResponse() *TaobaoTopEventSubscriptionQueryAPIResponse {
	return poolTaobaoTopEventSubscriptionQueryAPIResponse.Get().(*TaobaoTopEventSubscriptionQueryAPIResponse)
}

// ReleaseTaobaoTopEventSubscriptionQueryAPIResponse 将 TaobaoTopEventSubscriptionQueryAPIResponse 保存到 sync.Pool
func ReleaseTaobaoTopEventSubscriptionQueryAPIResponse(v *TaobaoTopEventSubscriptionQueryAPIResponse) {
	v.Reset()
	poolTaobaoTopEventSubscriptionQueryAPIResponse.Put(v)
}
