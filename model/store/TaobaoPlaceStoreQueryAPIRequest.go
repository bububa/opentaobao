package store

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoplacestorequeryAPIRequest 门店信息查询接口 API请求
// taobao.place.store.query
//
// 根据用户授权信息，获取用户的门店公开信息
type TaobaoplacestorequeryAPIRequest struct {
	model.Params
	// 业务code，用于区分业务
	_bizCode string
	// 业务外部id
	_outerId string
	// 门店id
	_storeId int64
}

// NewTaobaoplacestorequeryRequest 初始化TaobaoplacestorequeryAPIRequest对象
func NewTaobaoplacestorequeryRequest() *TaobaoplacestorequeryAPIRequest {
	return &TaobaoplacestorequeryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoplacestorequeryAPIRequest) GetApiMethodName() string {
	return "taobao.place.store.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoplacestorequeryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoplacestorequeryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBizCode is BizCode Setter
// 业务code，用于区分业务
func (r *TaobaoplacestorequeryAPIRequest) SetBizCode(_bizCode string) error {
	r._bizCode = _bizCode
	r.Set("biz_code", _bizCode)
	return nil
}

// GetBizCode BizCode Getter
func (r TaobaoplacestorequeryAPIRequest) GetBizCode() string {
	return r._bizCode
}

// SetOuterId is OuterId Setter
// 业务外部id
func (r *TaobaoplacestorequeryAPIRequest) SetOuterId(_outerId string) error {
	r._outerId = _outerId
	r.Set("outer_id", _outerId)
	return nil
}

// GetOuterId OuterId Getter
func (r TaobaoplacestorequeryAPIRequest) GetOuterId() string {
	return r._outerId
}

// SetStoreId is StoreId Setter
// 门店id
func (r *TaobaoplacestorequeryAPIRequest) SetStoreId(_storeId int64) error {
	r._storeId = _storeId
	r.Set("store_id", _storeId)
	return nil
}

// GetStoreId StoreId Getter
func (r TaobaoplacestorequeryAPIRequest) GetStoreId() int64 {
	return r._storeId
}
