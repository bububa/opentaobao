package alihouse

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseChangeStoreTypeAPIRequest 融合店迁移门店 API请求
// alibaba.alihouse.change.store.type
//
// 融合店迁移门店
type AlibabaAlihouseChangeStoreTypeAPIRequest struct {
	model.Params
	// 1
	_outerStoreId string
	// 1
	_outerCompanyId string
	// 1
	_storeType int64
	// 1
	_subType int64
}

// NewAlibabaAlihouseChangeStoreTypeRequest 初始化AlibabaAlihouseChangeStoreTypeAPIRequest对象
func NewAlibabaAlihouseChangeStoreTypeRequest() *AlibabaAlihouseChangeStoreTypeAPIRequest {
	return &AlibabaAlihouseChangeStoreTypeAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihouseChangeStoreTypeAPIRequest) Reset() {
	r._outerStoreId = ""
	r._outerCompanyId = ""
	r._storeType = 0
	r._subType = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseChangeStoreTypeAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.change.store.type"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseChangeStoreTypeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihouseChangeStoreTypeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOuterStoreId is OuterStoreId Setter
// 1
func (r *AlibabaAlihouseChangeStoreTypeAPIRequest) SetOuterStoreId(_outerStoreId string) error {
	r._outerStoreId = _outerStoreId
	r.Set("outer_store_id", _outerStoreId)
	return nil
}

// GetOuterStoreId OuterStoreId Getter
func (r AlibabaAlihouseChangeStoreTypeAPIRequest) GetOuterStoreId() string {
	return r._outerStoreId
}

// SetOuterCompanyId is OuterCompanyId Setter
// 1
func (r *AlibabaAlihouseChangeStoreTypeAPIRequest) SetOuterCompanyId(_outerCompanyId string) error {
	r._outerCompanyId = _outerCompanyId
	r.Set("outer_company_id", _outerCompanyId)
	return nil
}

// GetOuterCompanyId OuterCompanyId Getter
func (r AlibabaAlihouseChangeStoreTypeAPIRequest) GetOuterCompanyId() string {
	return r._outerCompanyId
}

// SetStoreType is StoreType Setter
// 1
func (r *AlibabaAlihouseChangeStoreTypeAPIRequest) SetStoreType(_storeType int64) error {
	r._storeType = _storeType
	r.Set("store_type", _storeType)
	return nil
}

// GetStoreType StoreType Getter
func (r AlibabaAlihouseChangeStoreTypeAPIRequest) GetStoreType() int64 {
	return r._storeType
}

// SetSubType is SubType Setter
// 1
func (r *AlibabaAlihouseChangeStoreTypeAPIRequest) SetSubType(_subType int64) error {
	r._subType = _subType
	r.Set("sub_type", _subType)
	return nil
}

// GetSubType SubType Getter
func (r AlibabaAlihouseChangeStoreTypeAPIRequest) GetSubType() int64 {
	return r._subType
}

var poolAlibabaAlihouseChangeStoreTypeAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihouseChangeStoreTypeRequest()
	},
}

// GetAlibabaAlihouseChangeStoreTypeRequest 从 sync.Pool 获取 AlibabaAlihouseChangeStoreTypeAPIRequest
func GetAlibabaAlihouseChangeStoreTypeAPIRequest() *AlibabaAlihouseChangeStoreTypeAPIRequest {
	return poolAlibabaAlihouseChangeStoreTypeAPIRequest.Get().(*AlibabaAlihouseChangeStoreTypeAPIRequest)
}

// ReleaseAlibabaAlihouseChangeStoreTypeAPIRequest 将 AlibabaAlihouseChangeStoreTypeAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihouseChangeStoreTypeAPIRequest(v *AlibabaAlihouseChangeStoreTypeAPIRequest) {
	v.Reset()
	poolAlibabaAlihouseChangeStoreTypeAPIRequest.Put(v)
}
