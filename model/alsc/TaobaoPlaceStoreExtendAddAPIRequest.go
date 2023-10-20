package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoplacestoreextendaddAPIRequest 新增门店扩展属性 API请求
// taobao.place.store.extend.add
//
// 新增授权用户的门店扩展属性
type TaobaoplacestoreextendaddAPIRequest struct {
	model.Params
	// 扩展信息
	_etv []ExtendTypeValueTopDto
	// 门店ID
	_storeId int64
}

// NewTaobaoplacestoreextendaddRequest 初始化TaobaoplacestoreextendaddAPIRequest对象
func NewTaobaoplacestoreextendaddRequest() *TaobaoplacestoreextendaddAPIRequest {
	return &TaobaoplacestoreextendaddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoplacestoreextendaddAPIRequest) GetApiMethodName() string {
	return "taobao.place.store.extend.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoplacestoreextendaddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoplacestoreextendaddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetEtv is Etv Setter
// 扩展信息
func (r *TaobaoplacestoreextendaddAPIRequest) SetEtv(_etv []ExtendTypeValueTopDto) error {
	r._etv = _etv
	r.Set("etv", _etv)
	return nil
}

// GetEtv Etv Getter
func (r TaobaoplacestoreextendaddAPIRequest) GetEtv() []ExtendTypeValueTopDto {
	return r._etv
}

// SetStoreId is StoreId Setter
// 门店ID
func (r *TaobaoplacestoreextendaddAPIRequest) SetStoreId(_storeId int64) error {
	r._storeId = _storeId
	r.Set("store_id", _storeId)
	return nil
}

// GetStoreId StoreId Getter
func (r TaobaoplacestoreextendaddAPIRequest) GetStoreId() int64 {
	return r._storeId
}
