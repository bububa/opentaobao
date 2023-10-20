package jym

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabajymsteamrefundauditAPIRequest 交易猫steam逆向单审核 API请求
// alibaba.jym.steam.refund.audit
//
// 交易猫steam逆向单审核
type AlibabajymsteamrefundauditAPIRequest struct {
	model.Params
	// 入参object
	_auditRefundOrderDto *AuditRefundOrderDto
}

// NewAlibabajymsteamrefundauditRequest 初始化AlibabajymsteamrefundauditAPIRequest对象
func NewAlibabajymsteamrefundauditRequest() *AlibabajymsteamrefundauditAPIRequest {
	return &AlibabajymsteamrefundauditAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabajymsteamrefundauditAPIRequest) GetApiMethodName() string {
	return "alibaba.jym.steam.refund.audit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabajymsteamrefundauditAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabajymsteamrefundauditAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAuditRefundOrderDto is AuditRefundOrderDto Setter
// 入参object
func (r *AlibabajymsteamrefundauditAPIRequest) SetAuditRefundOrderDto(_auditRefundOrderDto *AuditRefundOrderDto) error {
	r._auditRefundOrderDto = _auditRefundOrderDto
	r.Set("audit_refund_order_dto", _auditRefundOrderDto)
	return nil
}

// GetAuditRefundOrderDto AuditRefundOrderDto Getter
func (r AlibabajymsteamrefundauditAPIRequest) GetAuditRefundOrderDto() *AuditRefundOrderDto {
	return r._auditRefundOrderDto
}
