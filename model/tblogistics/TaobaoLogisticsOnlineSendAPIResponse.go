package tblogistics

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoLogisticsOnlineSendAPIResponse 在线订单发货处理（支持货到付款） API返回值
// taobao.logistics.online.send
//
// 用户调用该接口可实现在线订单发货（支持货到付款）&lt;br/&gt;调用该接口实现在线下单发货，有两种情况：&lt;br&gt;&lt;br/&gt;&lt;font color=&#39;red&#39;&gt;如果不输入运单号的情况：交易状态不会改变，需要调用taobao.logistics.online.confirm确认发货后交易状态才会变成卖家已发货。&lt;br&gt;&lt;br/&gt;如果输入运单号的情况发货：交易订单状态会直接变成卖家已发货 。&lt;/font&gt;
type TaobaoLogisticsOnlineSendAPIResponse struct {
	model.CommonResponse
	TaobaoLogisticsOnlineSendAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoLogisticsOnlineSendAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoLogisticsOnlineSendAPIResponseModel).Reset()
}

// TaobaoLogisticsOnlineSendAPIResponseModel is 在线订单发货处理（支持货到付款） 成功返回结果
type TaobaoLogisticsOnlineSendAPIResponseModel struct {
	XMLName xml.Name `xml:"logistics_online_send_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// de
	Shipping *Shipping `json:"shipping,omitempty" xml:"shipping,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoLogisticsOnlineSendAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Shipping = nil
}

var poolTaobaoLogisticsOnlineSendAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoLogisticsOnlineSendAPIResponse)
	},
}

// GetTaobaoLogisticsOnlineSendAPIResponse 从 sync.Pool 获取 TaobaoLogisticsOnlineSendAPIResponse
func GetTaobaoLogisticsOnlineSendAPIResponse() *TaobaoLogisticsOnlineSendAPIResponse {
	return poolTaobaoLogisticsOnlineSendAPIResponse.Get().(*TaobaoLogisticsOnlineSendAPIResponse)
}

// ReleaseTaobaoLogisticsOnlineSendAPIResponse 将 TaobaoLogisticsOnlineSendAPIResponse 保存到 sync.Pool
func ReleaseTaobaoLogisticsOnlineSendAPIResponse(v *TaobaoLogisticsOnlineSendAPIResponse) {
	v.Reset()
	poolTaobaoLogisticsOnlineSendAPIResponse.Put(v)
}
