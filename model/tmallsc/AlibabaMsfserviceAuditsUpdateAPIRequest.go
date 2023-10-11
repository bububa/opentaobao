package tmallsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMsfserviceAuditsUpdateAPIRequest 操作改约审批单 API请求
// alibaba.msfservice.audits.update
//
// 操作改约审批单
type AlibabaMsfserviceAuditsUpdateAPIRequest struct {
	model.Params
	// 审核接口入参
	_reserveAuditRequest *ReserveAuditRequest
}

// NewAlibabaMsfserviceAuditsUpdateRequest 初始化AlibabaMsfserviceAuditsUpdateAPIRequest对象
func NewAlibabaMsfserviceAuditsUpdateRequest() *AlibabaMsfserviceAuditsUpdateAPIRequest {
	return &AlibabaMsfserviceAuditsUpdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaMsfserviceAuditsUpdateAPIRequest) GetApiMethodName() string {
	return "alibaba.msfservice.audits.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaMsfserviceAuditsUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaMsfserviceAuditsUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetReserveAuditRequest is ReserveAuditRequest Setter
// 审核接口入参
func (r *AlibabaMsfserviceAuditsUpdateAPIRequest) SetReserveAuditRequest(_reserveAuditRequest *ReserveAuditRequest) error {
	r._reserveAuditRequest = _reserveAuditRequest
	r.Set("reserve_audit_request", _reserveAuditRequest)
	return nil
}

// GetReserveAuditRequest ReserveAuditRequest Getter
func (r AlibabaMsfserviceAuditsUpdateAPIRequest) GetReserveAuditRequest() *ReserveAuditRequest {
	return r._reserveAuditRequest
}
