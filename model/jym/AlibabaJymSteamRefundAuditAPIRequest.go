package jym

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaJymSteamRefundAuditAPIRequest 交易猫steam逆向单审核 API请求
// alibaba.jym.steam.refund.audit
//
// 交易猫steam逆向单审核
type AlibabaJymSteamRefundAuditAPIRequest struct {
	model.Params
	// 入参object
	_auditRefundOrderDto *AuditRefundOrderDto
}

// NewAlibabaJymSteamRefundAuditRequest 初始化AlibabaJymSteamRefundAuditAPIRequest对象
func NewAlibabaJymSteamRefundAuditRequest() *AlibabaJymSteamRefundAuditAPIRequest {
	return &AlibabaJymSteamRefundAuditAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaJymSteamRefundAuditAPIRequest) GetApiMethodName() string {
	return "alibaba.jym.steam.refund.audit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaJymSteamRefundAuditAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaJymSteamRefundAuditAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAuditRefundOrderDto is AuditRefundOrderDto Setter
// 入参object
func (r *AlibabaJymSteamRefundAuditAPIRequest) SetAuditRefundOrderDto(_auditRefundOrderDto *AuditRefundOrderDto) error {
	r._auditRefundOrderDto = _auditRefundOrderDto
	r.Set("audit_refund_order_dto", _auditRefundOrderDto)
	return nil
}

// GetAuditRefundOrderDto AuditRefundOrderDto Getter
func (r AlibabaJymSteamRefundAuditAPIRequest) GetAuditRefundOrderDto() *AuditRefundOrderDto {
	return r._auditRefundOrderDto
}
