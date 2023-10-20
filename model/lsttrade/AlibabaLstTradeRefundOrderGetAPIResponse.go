package lsttrade

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLstTradeRefundOrderGetAPIResponse 零售通退款订单查询 API返回值
// alibaba.lst.trade.refund.order.get
//
// 零售通退款订单查询
type AlibabaLstTradeRefundOrderGetAPIResponse struct {
	model.CommonResponse
	AlibabaLstTradeRefundOrderGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaLstTradeRefundOrderGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaLstTradeRefundOrderGetAPIResponseModel).Reset()
}

// AlibabaLstTradeRefundOrderGetAPIResponseModel is 零售通退款订单查询 成功返回结果
type AlibabaLstTradeRefundOrderGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_lst_trade_refund_order_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 子订单退款信息列表
	SubOrders []RefundSubOrderInfo `json:"sub_orders,omitempty" xml:"sub_orders>refund_sub_order_info,omitempty"`
	// 退款单id
	RefundId string `json:"refund_id,omitempty" xml:"refund_id,omitempty"`
	// 退款申请时间
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 退款状态，取值如下：等待卖家同意(waitselleragree),待买家修改(waitbuyermodify),等待买家退货(waitbuyersend),等待卖家确认收货(waitsellerreceive),退款成功(refundsuccess), 退款关闭(refundclose)
	RefundStatus string `json:"refund_status,omitempty" xml:"refund_status,omitempty"`
	// 退款完成时间
	GmtCompleted string `json:"gmt_completed,omitempty" xml:"gmt_completed,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaLstTradeRefundOrderGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.SubOrders = m.SubOrders[:0]
	m.RefundId = ""
	m.GmtCreate = ""
	m.RefundStatus = ""
	m.GmtCompleted = ""
}

var poolAlibabaLstTradeRefundOrderGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaLstTradeRefundOrderGetAPIResponse)
	},
}

// GetAlibabaLstTradeRefundOrderGetAPIResponse 从 sync.Pool 获取 AlibabaLstTradeRefundOrderGetAPIResponse
func GetAlibabaLstTradeRefundOrderGetAPIResponse() *AlibabaLstTradeRefundOrderGetAPIResponse {
	return poolAlibabaLstTradeRefundOrderGetAPIResponse.Get().(*AlibabaLstTradeRefundOrderGetAPIResponse)
}

// ReleaseAlibabaLstTradeRefundOrderGetAPIResponse 将 AlibabaLstTradeRefundOrderGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaLstTradeRefundOrderGetAPIResponse(v *AlibabaLstTradeRefundOrderGetAPIResponse) {
	v.Reset()
	poolAlibabaLstTradeRefundOrderGetAPIResponse.Put(v)
}
