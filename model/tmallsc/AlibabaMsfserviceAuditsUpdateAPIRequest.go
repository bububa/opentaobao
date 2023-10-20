package tmallsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabamsfserviceauditsupdateAPIRequest 操作改约审批单 API请求
// alibaba.msfservice.audits.update
//
// 操作改约审批单
type AlibabamsfserviceauditsupdateAPIRequest struct {
	model.Params
	// 审核接口入参
	_reserveAuditRequest *ReserveAuditRequest
}

// NewAlibabamsfserviceauditsupdateRequest 初始化AlibabamsfserviceauditsupdateAPIRequest对象
func NewAlibabamsfserviceauditsupdateRequest() *AlibabamsfserviceauditsupdateAPIRequest {
	return &AlibabamsfserviceauditsupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabamsfserviceauditsupdateAPIRequest) GetApiMethodName() string {
	return "alibaba.msfservice.audits.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabamsfserviceauditsupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabamsfserviceauditsupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetReserveAuditRequest is ReserveAuditRequest Setter
// 审核接口入参
func (r *AlibabamsfserviceauditsupdateAPIRequest) SetReserveAuditRequest(_reserveAuditRequest *ReserveAuditRequest) error {
	r._reserveAuditRequest = _reserveAuditRequest
	r.Set("reserve_audit_request", _reserveAuditRequest)
	return nil
}

// GetReserveAuditRequest ReserveAuditRequest Getter
func (r AlibabamsfserviceauditsupdateAPIRequest) GetReserveAuditRequest() *ReserveAuditRequest {
	return r._reserveAuditRequest
}
