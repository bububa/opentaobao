package fenxiao

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFenxiaoOrderConfirmPaidAPIResponse 确认收款 API返回值
// taobao.fenxiao.order.confirm.paid
//
// 供应商确认收款（非支付宝交易）。
type TaobaoFenxiaoOrderConfirmPaidAPIResponse struct {
	model.CommonResponse
	TaobaoFenxiaoOrderConfirmPaidAPIResponseModel
}

// TaobaoFenxiaoOrderConfirmPaidAPIResponseModel is 确认收款 成功返回结果
type TaobaoFenxiaoOrderConfirmPaidAPIResponseModel struct {
	XMLName xml.Name `xml:"fenxiao_order_confirm_paid_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 确认结果成功与否
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
