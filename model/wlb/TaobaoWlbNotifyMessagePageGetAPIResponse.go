package wlb

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWlbNotifyMessagePageGetAPIResponse 物流宝通知消息查询接口 API返回值
// taobao.wlb.notify.message.page.get
//
// 物流宝提供的消息通知查询接口，消息内容包括;出入库单不一致消息，取消订单成功消息，盘点单消息
type TaobaoWlbNotifyMessagePageGetAPIResponse struct {
	model.CommonResponse
	TaobaoWlbNotifyMessagePageGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoWlbNotifyMessagePageGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoWlbNotifyMessagePageGetAPIResponseModel).Reset()
}

// TaobaoWlbNotifyMessagePageGetAPIResponseModel is 物流宝通知消息查询接口 成功返回结果
type TaobaoWlbNotifyMessagePageGetAPIResponseModel struct {
	XMLName xml.Name `xml:"wlb_notify_message_page_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 通道消息
	WlbMessages []WlbMessage `json:"wlb_messages,omitempty" xml:"wlb_messages>wlb_message,omitempty"`
	// 2000-01-01 00:00:00
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoWlbNotifyMessagePageGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.WlbMessages = m.WlbMessages[:0]
	m.TotalCount = 0
}

var poolTaobaoWlbNotifyMessagePageGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoWlbNotifyMessagePageGetAPIResponse)
	},
}

// GetTaobaoWlbNotifyMessagePageGetAPIResponse 从 sync.Pool 获取 TaobaoWlbNotifyMessagePageGetAPIResponse
func GetTaobaoWlbNotifyMessagePageGetAPIResponse() *TaobaoWlbNotifyMessagePageGetAPIResponse {
	return poolTaobaoWlbNotifyMessagePageGetAPIResponse.Get().(*TaobaoWlbNotifyMessagePageGetAPIResponse)
}

// ReleaseTaobaoWlbNotifyMessagePageGetAPIResponse 将 TaobaoWlbNotifyMessagePageGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoWlbNotifyMessagePageGetAPIResponse(v *TaobaoWlbNotifyMessagePageGetAPIResponse) {
	v.Reset()
	poolTaobaoWlbNotifyMessagePageGetAPIResponse.Put(v)
}
