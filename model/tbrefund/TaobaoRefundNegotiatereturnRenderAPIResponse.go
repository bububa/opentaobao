package tbrefund

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoRefundNegotiatereturnRenderAPIResponse 协商退货退款渲染 API返回值
// taobao.refund.negotiatereturn.render
//
// 协商退货退款渲染
type TaobaoRefundNegotiatereturnRenderAPIResponse struct {
	model.CommonResponse
	TaobaoRefundNegotiatereturnRenderAPIResponseModel
}

// TaobaoRefundNegotiatereturnRenderAPIResponseModel is 协商退货退款渲染 成功返回结果
type TaobaoRefundNegotiatereturnRenderAPIResponseModel struct {
	XMLName xml.Name `xml:"refund_negotiatereturn_render_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 可以协商的退款原因列表
	ReasonList []Reason `json:"reason_list,omitempty" xml:"reason_list>reason,omitempty"`
	// 卖家退货地址列表
	AddressList []Address `json:"address_list,omitempty" xml:"address_list>address,omitempty"`
	// 申请协商提示文案
	ApplyTips string `json:"apply_tips,omitempty" xml:"apply_tips,omitempty"`
	// 退款版本号
	RefundVersion int64 `json:"refund_version,omitempty" xml:"refund_version,omitempty"`
	// 可以协商的最大退款金额
	MaxRefundFee *MaxRefundFee `json:"max_refund_fee,omitempty" xml:"max_refund_fee,omitempty"`
}
