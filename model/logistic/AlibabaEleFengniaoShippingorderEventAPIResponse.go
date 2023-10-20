package logistic

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEleFengniaoShippingorderEventAPIResponse 查询运单事件信息 API返回值
// alibaba.ele.fengniao.shippingorder.event
//
// 查询运单事件信息
type AlibabaEleFengniaoShippingorderEventAPIResponse struct {
	model.CommonResponse
	AlibabaEleFengniaoShippingorderEventAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaEleFengniaoShippingorderEventAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaEleFengniaoShippingorderEventAPIResponseModel).Reset()
}

// AlibabaEleFengniaoShippingorderEventAPIResponseModel is 查询运单事件信息 成功返回结果
type AlibabaEleFengniaoShippingorderEventAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ele_fengniao_shippingorder_event_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// shippingOrderEvents
	ShippingOrderEvents []ShippingOrderEvent `json:"shipping_order_events,omitempty" xml:"shipping_order_events>shipping_order_event,omitempty"`
	// MERCHANT_CANCEL:商家取消,DELIVERY_TIMEOUT:配送超时，系统标记异常
	ShippingRemarkCode string `json:"shipping_remark_code,omitempty" xml:"shipping_remark_code,omitempty"`
	// 终态时间
	FinishAt int64 `json:"finish_at,omitempty" xml:"finish_at,omitempty"`
	// 骑手预计送达时间
	PredictDeliveryAt int64 `json:"predict_delivery_at,omitempty" xml:"predict_delivery_at,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaEleFengniaoShippingorderEventAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ShippingOrderEvents = m.ShippingOrderEvents[:0]
	m.ShippingRemarkCode = ""
	m.FinishAt = 0
	m.PredictDeliveryAt = 0
}

var poolAlibabaEleFengniaoShippingorderEventAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaEleFengniaoShippingorderEventAPIResponse)
	},
}

// GetAlibabaEleFengniaoShippingorderEventAPIResponse 从 sync.Pool 获取 AlibabaEleFengniaoShippingorderEventAPIResponse
func GetAlibabaEleFengniaoShippingorderEventAPIResponse() *AlibabaEleFengniaoShippingorderEventAPIResponse {
	return poolAlibabaEleFengniaoShippingorderEventAPIResponse.Get().(*AlibabaEleFengniaoShippingorderEventAPIResponse)
}

// ReleaseAlibabaEleFengniaoShippingorderEventAPIResponse 将 AlibabaEleFengniaoShippingorderEventAPIResponse 保存到 sync.Pool
func ReleaseAlibabaEleFengniaoShippingorderEventAPIResponse(v *AlibabaEleFengniaoShippingorderEventAPIResponse) {
	v.Reset()
	poolAlibabaEleFengniaoShippingorderEventAPIResponse.Put(v)
}
