package bus

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaobusordersetAPIResponse 汽车票下单接口 API返回值
// taobao.bus.order.set
//
// 提供给汽车票商家进行下单
type TaobaobusordersetAPIResponse struct {
	model.CommonResponse
	TaobaobusordersetAPIResponseModel
}

// TaobaobusordersetAPIResponseModel is 汽车票下单接口 成功返回结果
type TaobaobusordersetAPIResponseModel struct {
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
