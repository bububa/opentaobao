package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPlaceStoreExtendAddAPIRequest 新增门店扩展属性 API请求
// taobao.place.store.extend.add
//
// 新增授权用户的门店扩展属性
type TaobaoPlaceStoreExtendAddAPIRequest struct {
	model.Params
	// 扩展信息
	_etv []ExtendTypeValueTopDto
	// 门店ID
	_storeId int64
}

// NewTaobaoPlaceStoreExtendAddRequest 初始化TaobaoPlaceStoreExtendAddAPIRequest对象
func NewTaobaoPlaceStoreExtendAddRequest() *TaobaoPlaceStoreExtendAddAPIRequest {
	return &TaobaoPlaceStoreExtendAddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoPlaceStoreExtendAddAPIRequest) GetApiMethodName() string {
	return "taobao.place.store.extend.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoPlaceStoreExtendAddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoPlaceStoreExtendAddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetEtv is Etv Setter
// 扩展信息
func (r *TaobaoPlaceStoreExtendAddAPIRequest) SetEtv(_etv []ExtendTypeValueTopDto) error {
	r._etv = _etv
	r.Set("etv", _etv)
	return nil
}

// GetEtv Etv Getter
func (r TaobaoPlaceStoreExtendAddAPIRequest) GetEtv() []ExtendTypeValueTopDto {
	return r._etv
}

// SetStoreId is StoreId Setter
// 门店ID
func (r *TaobaoPlaceStoreExtendAddAPIRequest) SetStoreId(_storeId int64) error {
	r._storeId = _storeId
	r.Set("store_id", _storeId)
	return nil
}

// GetStoreId StoreId Getter
func (r TaobaoPlaceStoreExtendAddAPIRequest) GetStoreId() int64 {
	return r._storeId
}
