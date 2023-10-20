package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihousechangestoretypeAPIRequest 融合店迁移门店 API请求
// alibaba.alihouse.change.store.type
//
// 融合店迁移门店
type AlibabaalihousechangestoretypeAPIRequest struct {
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

// NewAlibabaalihousechangestoretypeRequest 初始化AlibabaalihousechangestoretypeAPIRequest对象
func NewAlibabaalihousechangestoretypeRequest() *AlibabaalihousechangestoretypeAPIRequest {
	return &AlibabaalihousechangestoretypeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihousechangestoretypeAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.change.store.type"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihousechangestoretypeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihousechangestoretypeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOuterStoreId is OuterStoreId Setter
// 1
func (r *AlibabaalihousechangestoretypeAPIRequest) SetOuterStoreId(_outerStoreId string) error {
	r._outerStoreId = _outerStoreId
	r.Set("outer_store_id", _outerStoreId)
	return nil
}

// GetOuterStoreId OuterStoreId Getter
func (r AlibabaalihousechangestoretypeAPIRequest) GetOuterStoreId() string {
	return r._outerStoreId
}

// SetOuterCompanyId is OuterCompanyId Setter
// 1
func (r *AlibabaalihousechangestoretypeAPIRequest) SetOuterCompanyId(_outerCompanyId string) error {
	r._outerCompanyId = _outerCompanyId
	r.Set("outer_company_id", _outerCompanyId)
	return nil
}

// GetOuterCompanyId OuterCompanyId Getter
func (r AlibabaalihousechangestoretypeAPIRequest) GetOuterCompanyId() string {
	return r._outerCompanyId
}

// SetStoreType is StoreType Setter
// 1
func (r *AlibabaalihousechangestoretypeAPIRequest) SetStoreType(_storeType int64) error {
	r._storeType = _storeType
	r.Set("store_type", _storeType)
	return nil
}

// GetStoreType StoreType Getter
func (r AlibabaalihousechangestoretypeAPIRequest) GetStoreType() int64 {
	return r._storeType
}

// SetSubType is SubType Setter
// 1
func (r *AlibabaalihousechangestoretypeAPIRequest) SetSubType(_subType int64) error {
	r._subType = _subType
	r.Set("sub_type", _subType)
	return nil
}

// GetSubType SubType Getter
func (r AlibabaalihousechangestoretypeAPIRequest) GetSubType() int64 {
	return r._subType
}
