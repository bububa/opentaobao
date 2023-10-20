package logistic

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoLogisticsOfflineSendAPIResponse 自己联系物流（线下物流）发货 API返回值
// taobao.logistics.offline.send
//
// 用户调用该接口可实现自己联系发货（线下物流），使用该接口发货，交易订单状态会直接变成卖家已发货。不支持货到付款、在线下单类型的订单。
type TaobaoLogisticsOfflineSendAPIResponse struct {
	model.CommonResponse
	TaobaoLogisticsOfflineSendAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoLogisticsOfflineSendAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoLogisticsOfflineSendAPIResponseModel).Reset()
}

// TaobaoLogisticsOfflineSendAPIResponseModel is 自己联系物流（线下物流）发货 成功返回结果
type TaobaoLogisticsOfflineSendAPIResponseModel struct {
	XMLName xml.Name `xml:"logistics_offline_send_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 自己联系的调用结果
	Shipping *Shipping `json:"shipping,omitempty" xml:"shipping,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoLogisticsOfflineSendAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Shipping = nil
}

var poolTaobaoLogisticsOfflineSendAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoLogisticsOfflineSendAPIResponse)
	},
}

// GetTaobaoLogisticsOfflineSendAPIResponse 从 sync.Pool 获取 TaobaoLogisticsOfflineSendAPIResponse
func GetTaobaoLogisticsOfflineSendAPIResponse() *TaobaoLogisticsOfflineSendAPIResponse {
	return poolTaobaoLogisticsOfflineSendAPIResponse.Get().(*TaobaoLogisticsOfflineSendAPIResponse)
}

// ReleaseTaobaoLogisticsOfflineSendAPIResponse 将 TaobaoLogisticsOfflineSendAPIResponse 保存到 sync.Pool
func ReleaseTaobaoLogisticsOfflineSendAPIResponse(v *TaobaoLogisticsOfflineSendAPIResponse) {
	v.Reset()
	poolTaobaoLogisticsOfflineSendAPIResponse.Put(v)
}
