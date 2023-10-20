package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallalihousetradecouponorderauditrefundAPIRequest ETC审核电商券退款 API请求
// tmall.alihouse.trade.coupon.order.audit.refund
//
// ETC审核电商券退款
type TmallalihousetradecouponorderauditrefundAPIRequest struct {
	model.Params
	// 审核退款
	_auditOrderDto *AuditOrderDto
}

// NewTmallalihousetradecouponorderauditrefundRequest 初始化TmallalihousetradecouponorderauditrefundAPIRequest对象
func NewTmallalihousetradecouponorderauditrefundRequest() *TmallalihousetradecouponorderauditrefundAPIRequest {
	return &TmallalihousetradecouponorderauditrefundAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallalihousetradecouponorderauditrefundAPIRequest) GetApiMethodName() string {
	return "tmall.alihouse.trade.coupon.order.audit.refund"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallalihousetradecouponorderauditrefundAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallalihousetradecouponorderauditrefundAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAuditOrderDto is AuditOrderDto Setter
// 审核退款
func (r *TmallalihousetradecouponorderauditrefundAPIRequest) SetAuditOrderDto(_auditOrderDto *AuditOrderDto) error {
	r._auditOrderDto = _auditOrderDto
	r.Set("audit_order_dto", _auditOrderDto)
	return nil
}

// GetAuditOrderDto AuditOrderDto Getter
func (r TmallalihousetradecouponorderauditrefundAPIRequest) GetAuditOrderDto() *AuditOrderDto {
	return r._auditOrderDto
}
