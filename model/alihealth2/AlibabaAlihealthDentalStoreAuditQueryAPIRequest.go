package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdentalstoreauditqueryAPIRequest ISV查询门店审核状态 API请求
// alibaba.alihealth.dental.store.audit.query
//
// ISV查询门店审核状态
type AlibabaalihealthdentalstoreauditqueryAPIRequest struct {
	model.Params
	// 审核ID列表
	_storeAuditIds []string
}

// NewAlibabaalihealthdentalstoreauditqueryRequest 初始化AlibabaalihealthdentalstoreauditqueryAPIRequest对象
func NewAlibabaalihealthdentalstoreauditqueryRequest() *AlibabaalihealthdentalstoreauditqueryAPIRequest {
	return &AlibabaalihealthdentalstoreauditqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthdentalstoreauditqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.dental.store.audit.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthdentalstoreauditqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthdentalstoreauditqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStoreAuditIds is StoreAuditIds Setter
// 审核ID列表
func (r *AlibabaalihealthdentalstoreauditqueryAPIRequest) SetStoreAuditIds(_storeAuditIds []string) error {
	r._storeAuditIds = _storeAuditIds
	r.Set("store_audit_ids", _storeAuditIds)
	return nil
}

// GetStoreAuditIds StoreAuditIds Getter
func (r AlibabaalihealthdentalstoreauditqueryAPIRequest) GetStoreAuditIds() []string {
	return r._storeAuditIds
}
