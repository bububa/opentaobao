package alihouse

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallAlihouseTradeCouponRefundOrderQueryAPIResponse 查询电商券履约退款单 API返回值
// tmall.alihouse.trade.coupon.refund.order.query
//
// 查询电商券履约退款单
type TmallAlihouseTradeCouponRefundOrderQueryAPIResponse struct {
	model.CommonResponse
	TmallAlihouseTradeCouponRefundOrderQueryAPIResponseModel
}

// Reset 清空结构体
func (m *TmallAlihouseTradeCouponRefundOrderQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallAlihouseTradeCouponRefundOrderQueryAPIResponseModel).Reset()
}

// TmallAlihouseTradeCouponRefundOrderQueryAPIResponseModel is 查询电商券履约退款单 成功返回结果
type TmallAlihouseTradeCouponRefundOrderQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_alihouse_trade_coupon_refund_order_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *TmallAlihouseTradeCouponRefundOrderQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallAlihouseTradeCouponRefundOrderQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallAlihouseTradeCouponRefundOrderQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallAlihouseTradeCouponRefundOrderQueryAPIResponse)
	},
}

// GetTmallAlihouseTradeCouponRefundOrderQueryAPIResponse 从 sync.Pool 获取 TmallAlihouseTradeCouponRefundOrderQueryAPIResponse
func GetTmallAlihouseTradeCouponRefundOrderQueryAPIResponse() *TmallAlihouseTradeCouponRefundOrderQueryAPIResponse {
	return poolTmallAlihouseTradeCouponRefundOrderQueryAPIResponse.Get().(*TmallAlihouseTradeCouponRefundOrderQueryAPIResponse)
}

// ReleaseTmallAlihouseTradeCouponRefundOrderQueryAPIResponse 将 TmallAlihouseTradeCouponRefundOrderQueryAPIResponse 保存到 sync.Pool
func ReleaseTmallAlihouseTradeCouponRefundOrderQueryAPIResponse(v *TmallAlihouseTradeCouponRefundOrderQueryAPIResponse) {
	v.Reset()
	poolTmallAlihouseTradeCouponRefundOrderQueryAPIResponse.Put(v)
}
