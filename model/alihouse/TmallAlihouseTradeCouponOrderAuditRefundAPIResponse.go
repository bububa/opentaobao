package alihouse

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallAlihouseTradeCouponOrderAuditRefundAPIResponse ETC审核电商券退款 API返回值
// tmall.alihouse.trade.coupon.order.audit.refund
//
// ETC审核电商券退款
type TmallAlihouseTradeCouponOrderAuditRefundAPIResponse struct {
	model.CommonResponse
	TmallAlihouseTradeCouponOrderAuditRefundAPIResponseModel
}

// Reset 清空结构体
func (m *TmallAlihouseTradeCouponOrderAuditRefundAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallAlihouseTradeCouponOrderAuditRefundAPIResponseModel).Reset()
}

// TmallAlihouseTradeCouponOrderAuditRefundAPIResponseModel is ETC审核电商券退款 成功返回结果
type TmallAlihouseTradeCouponOrderAuditRefundAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_alihouse_trade_coupon_order_audit_refund_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *TmallAlihouseTradeCouponOrderAuditRefundResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallAlihouseTradeCouponOrderAuditRefundAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallAlihouseTradeCouponOrderAuditRefundAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallAlihouseTradeCouponOrderAuditRefundAPIResponse)
	},
}

// GetTmallAlihouseTradeCouponOrderAuditRefundAPIResponse 从 sync.Pool 获取 TmallAlihouseTradeCouponOrderAuditRefundAPIResponse
func GetTmallAlihouseTradeCouponOrderAuditRefundAPIResponse() *TmallAlihouseTradeCouponOrderAuditRefundAPIResponse {
	return poolTmallAlihouseTradeCouponOrderAuditRefundAPIResponse.Get().(*TmallAlihouseTradeCouponOrderAuditRefundAPIResponse)
}

// ReleaseTmallAlihouseTradeCouponOrderAuditRefundAPIResponse 将 TmallAlihouseTradeCouponOrderAuditRefundAPIResponse 保存到 sync.Pool
func ReleaseTmallAlihouseTradeCouponOrderAuditRefundAPIResponse(v *TmallAlihouseTradeCouponOrderAuditRefundAPIResponse) {
	v.Reset()
	poolTmallAlihouseTradeCouponOrderAuditRefundAPIResponse.Put(v)
}
