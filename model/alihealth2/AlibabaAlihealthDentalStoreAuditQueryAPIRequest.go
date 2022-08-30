package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDentalStoreAuditQueryAPIRequest ISV查询门店审核状态 API请求
// alibaba.alihealth.dental.store.audit.query
//
// ISV查询门店审核状态
type AlibabaAlihealthDentalStoreAuditQueryAPIRequest struct {
	model.Params
	// 审核ID列表
	_storeAuditIds []string
}

// NewAlibabaAlihealthDentalStoreAuditQueryRequest 初始化AlibabaAlihealthDentalStoreAuditQueryAPIRequest对象
func NewAlibabaAlihealthDentalStoreAuditQueryRequest() *AlibabaAlihealthDentalStoreAuditQueryAPIRequest {
	return &AlibabaAlihealthDentalStoreAuditQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDentalStoreAuditQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.dental.store.audit.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDentalStoreAuditQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetStoreAuditIds is StoreAuditIds Setter
// 审核ID列表
func (r *AlibabaAlihealthDentalStoreAuditQueryAPIRequest) SetStoreAuditIds(_storeAuditIds []string) error {
	r._storeAuditIds = _storeAuditIds
	r.Set("store_audit_ids", _storeAuditIds)
	return nil
}

// GetStoreAuditIds StoreAuditIds Getter
func (r AlibabaAlihealthDentalStoreAuditQueryAPIRequest) GetStoreAuditIds() []string {
	return r._storeAuditIds
}
