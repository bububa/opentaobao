package alihouse

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallAlihouseTradeCouponOrderAuditRefundAPIRequest ETC审核电商券退款 API请求
// tmall.alihouse.trade.coupon.order.audit.refund
//
// ETC审核电商券退款
type TmallAlihouseTradeCouponOrderAuditRefundAPIRequest struct {
	model.Params
	// 审核退款
	_auditOrderDto *AuditOrderDto
}

// NewTmallAlihouseTradeCouponOrderAuditRefundRequest 初始化TmallAlihouseTradeCouponOrderAuditRefundAPIRequest对象
func NewTmallAlihouseTradeCouponOrderAuditRefundRequest() *TmallAlihouseTradeCouponOrderAuditRefundAPIRequest {
	return &TmallAlihouseTradeCouponOrderAuditRefundAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallAlihouseTradeCouponOrderAuditRefundAPIRequest) Reset() {
	r._auditOrderDto = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallAlihouseTradeCouponOrderAuditRefundAPIRequest) GetApiMethodName() string {
	return "tmall.alihouse.trade.coupon.order.audit.refund"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallAlihouseTradeCouponOrderAuditRefundAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallAlihouseTradeCouponOrderAuditRefundAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAuditOrderDto is AuditOrderDto Setter
// 审核退款
func (r *TmallAlihouseTradeCouponOrderAuditRefundAPIRequest) SetAuditOrderDto(_auditOrderDto *AuditOrderDto) error {
	r._auditOrderDto = _auditOrderDto
	r.Set("audit_order_dto", _auditOrderDto)
	return nil
}

// GetAuditOrderDto AuditOrderDto Getter
func (r TmallAlihouseTradeCouponOrderAuditRefundAPIRequest) GetAuditOrderDto() *AuditOrderDto {
	return r._auditOrderDto
}

var poolTmallAlihouseTradeCouponOrderAuditRefundAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallAlihouseTradeCouponOrderAuditRefundRequest()
	},
}

// GetTmallAlihouseTradeCouponOrderAuditRefundRequest 从 sync.Pool 获取 TmallAlihouseTradeCouponOrderAuditRefundAPIRequest
func GetTmallAlihouseTradeCouponOrderAuditRefundAPIRequest() *TmallAlihouseTradeCouponOrderAuditRefundAPIRequest {
	return poolTmallAlihouseTradeCouponOrderAuditRefundAPIRequest.Get().(*TmallAlihouseTradeCouponOrderAuditRefundAPIRequest)
}

// ReleaseTmallAlihouseTradeCouponOrderAuditRefundAPIRequest 将 TmallAlihouseTradeCouponOrderAuditRefundAPIRequest 放入 sync.Pool
func ReleaseTmallAlihouseTradeCouponOrderAuditRefundAPIRequest(v *TmallAlihouseTradeCouponOrderAuditRefundAPIRequest) {
	v.Reset()
	poolTmallAlihouseTradeCouponOrderAuditRefundAPIRequest.Put(v)
}
