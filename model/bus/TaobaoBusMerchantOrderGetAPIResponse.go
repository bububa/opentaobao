package bus

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaobusmerchantordergetAPIResponse 商家侧查询订单详情 API返回值
// taobao.bus.merchant.order.get
//
// 商家侧查询订单详情
type TaobaobusmerchantordergetAPIResponse struct {
	model.CommonResponse
	TaobaobusmerchantordergetAPIResponseModel
}

// TaobaobusmerchantordergetAPIResponseModel is 商家侧查询订单详情 成功返回结果
type TaobaobusmerchantordergetAPIResponseModel struct {
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
