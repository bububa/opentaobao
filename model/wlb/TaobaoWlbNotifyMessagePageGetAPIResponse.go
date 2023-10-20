package wlb

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaowlbnotifymessagepagegetAPIResponse 物流宝通知消息查询接口 API返回值
// taobao.wlb.notify.message.page.get
//
// 物流宝提供的消息通知查询接口，消息内容包括;出入库单不一致消息，取消订单成功消息，盘点单消息
type TaobaowlbnotifymessagepagegetAPIResponse struct {
	model.CommonResponse
	TaobaowlbnotifymessagepagegetAPIResponseModel
}

// TaobaowlbnotifymessagepagegetAPIResponseModel is 物流宝通知消息查询接口 成功返回结果
type TaobaowlbnotifymessagepagegetAPIResponseModel struct {
	XMLName xml.Name `xml:"wlb_notify_message_page_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 通道消息
	WlbMessages []WlbMessage `json:"wlb_messages,omitempty" xml:"wlb_messages>wlb_message,omitempty"`
	// 2000-01-01 00:00:00
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
}
