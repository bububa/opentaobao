package alihouse

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallAlihouseTradeCouponOrderAuditRefundAPIRequest) GetApiMethodName() string {
	return "tmall.alihouse.trade.coupon.order.audit.refund"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallAlihouseTradeCouponOrderAuditRefundAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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
