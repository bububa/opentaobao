package bus

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBusOrderSetAPIResponse 汽车票下单接口 API返回值
// taobao.bus.order.set
//
// 提供给汽车票商家进行下单
type TaobaoBusOrderSetAPIResponse struct {
	model.CommonResponse
	TaobaoBusOrderSetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoBusOrderSetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoBusOrderSetAPIResponseModel).Reset()
}

// TaobaoBusOrderSetAPIResponseModel is 汽车票下单接口 成功返回结果
type TaobaoBusOrderSetAPIResponseModel struct {
	XMLName xml.Name `xml:"bus_order_set_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 支付宝交易流水号
	AliPayTradeId string `json:"ali_pay_trade_id,omitempty" xml:"ali_pay_trade_id,omitempty"`
	// 阿里订单号
	AlitripOrderId string `json:"alitrip_order_id,omitempty" xml:"alitrip_order_id,omitempty"`
	// 错误代码
	ErrorCode1 string `json:"error_code_1,omitempty" xml:"error_code_1,omitempty"`
	// 错误描述
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 是否下单成功
	Issuccess bool `json:"issuccess,omitempty" xml:"issuccess,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoBusOrderSetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.AliPayTradeId = ""
	m.AlitripOrderId = ""
	m.ErrorCode1 = ""
	m.ErrorMsg = ""
	m.Issuccess = false
}

var poolTaobaoBusOrderSetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoBusOrderSetAPIResponse)
	},
}

// GetTaobaoBusOrderSetAPIResponse 从 sync.Pool 获取 TaobaoBusOrderSetAPIResponse
func GetTaobaoBusOrderSetAPIResponse() *TaobaoBusOrderSetAPIResponse {
	return poolTaobaoBusOrderSetAPIResponse.Get().(*TaobaoBusOrderSetAPIResponse)
}

// ReleaseTaobaoBusOrderSetAPIResponse 将 TaobaoBusOrderSetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoBusOrderSetAPIResponse(v *TaobaoBusOrderSetAPIResponse) {
	v.Reset()
	poolTaobaoBusOrderSetAPIResponse.Put(v)
}
