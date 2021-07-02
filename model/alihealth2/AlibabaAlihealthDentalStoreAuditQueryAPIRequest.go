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
	_storeAuditIds []int64
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

// Set is StoreAuditIds Setter
// 审核ID列表
func (r *AlibabaAlihealthDentalStoreAuditQueryAPIRequest) SetStoreAuditIds(_storeAuditIds []int64) error {
	r._storeAuditIds = _storeAuditIds
	r.Set("store_audit_ids", _storeAuditIds)
	return nil
}

// Get StoreAuditIds Getter
func (r AlibabaAlihealthDentalStoreAuditQueryAPIRequest) GetStoreAuditIds() []int64 {
	return r._storeAuditIds
}
