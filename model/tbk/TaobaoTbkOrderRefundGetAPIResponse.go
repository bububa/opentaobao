package tbk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaotbkorderrefundgetAPIResponse 淘宝客-推广者-全量维权退款订单查询 API返回值
// taobao.tbk.order.refund.get
//
// 淘宝客账户下全量维权退款订单查询
type TaobaotbkorderrefundgetAPIResponse struct {
	model.CommonResponse
	TaobaotbkorderrefundgetAPIResponseModel
}

// TaobaotbkorderrefundgetAPIResponseModel is 淘宝客-推广者-全量维权退款订单查询 成功返回结果
type TaobaotbkorderrefundgetAPIResponseModel struct {
	XMLName xml.Name `xml:"tbk_order_refund_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 业务错误信息
	Bizerrordesc string `json:"biz_error_desc,omitempty" xml:"biz_error_desc,omitempty"`
	// 返回信息
	Resultmsg string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 返回数据
	Data *OrderPage `json:"data,omitempty" xml:"data,omitempty"`
	// 接口返回值信息，跟rpc架构保持一致
	Resultcode int64 `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 业务错误码 101, 102,103
	Bizerrorcode int64 `json:"biz_error_code,omitempty" xml:"biz_error_code,omitempty"`
}
