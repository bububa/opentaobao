package bus

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBusMerchantOrderGetAPIResponse 商家侧查询订单详情 API返回值
// taobao.bus.merchant.order.get
//
// 商家侧查询订单详情
type TaobaoBusMerchantOrderGetAPIResponse struct {
	model.CommonResponse
	TaobaoBusMerchantOrderGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoBusMerchantOrderGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoBusMerchantOrderGetAPIResponseModel).Reset()
}

// TaobaoBusMerchantOrderGetAPIResponseModel is 商家侧查询订单详情 成功返回结果
type TaobaoBusMerchantOrderGetAPIResponseModel struct {
	XMLName xml.Name `xml:"bus_merchant_order_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 退票订单列表
	RefundOrderInfos []MerchantBusRefundOrderInfo `json:"refund_order_infos,omitempty" xml:"refund_order_infos>merchant_bus_refund_order_info,omitempty"`
	// 订单信息列表
	OrderInfos []MerchantBusOrderInfo `json:"order_infos,omitempty" xml:"order_infos>merchant_bus_order_info,omitempty"`
	// 错误码
	ErrorMsgCode string `json:"error_msg_code,omitempty" xml:"error_msg_code,omitempty"`
	// 错误描述
	ErrorMsgDesc string `json:"error_msg_desc,omitempty" xml:"error_msg_desc,omitempty"`
	// 根据总数进行分页查询
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
	// 业务接口是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoBusMerchantOrderGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.RefundOrderInfos = m.RefundOrderInfos[:0]
	m.OrderInfos = m.OrderInfos[:0]
	m.ErrorMsgCode = ""
	m.ErrorMsgDesc = ""
	m.TotalCount = 0
	m.IsSuccess = false
}

var poolTaobaoBusMerchantOrderGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoBusMerchantOrderGetAPIResponse)
	},
}

// GetTaobaoBusMerchantOrderGetAPIResponse 从 sync.Pool 获取 TaobaoBusMerchantOrderGetAPIResponse
func GetTaobaoBusMerchantOrderGetAPIResponse() *TaobaoBusMerchantOrderGetAPIResponse {
	return poolTaobaoBusMerchantOrderGetAPIResponse.Get().(*TaobaoBusMerchantOrderGetAPIResponse)
}

// ReleaseTaobaoBusMerchantOrderGetAPIResponse 将 TaobaoBusMerchantOrderGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoBusMerchantOrderGetAPIResponse(v *TaobaoBusMerchantOrderGetAPIResponse) {
	v.Reset()
	poolTaobaoBusMerchantOrderGetAPIResponse.Put(v)
}
