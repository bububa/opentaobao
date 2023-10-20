package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoqimenstorequeryAPIRequest 门店信息查询接口 API请求
// taobao.qimen.store.query
//
// 商家在ERP等系统中调用该接口，查询门店相关信息
type TaobaoqimenstorequeryAPIRequest struct {
	model.Params
	// 已分配的线上门店ID
	_storeId int64
}

// NewTaobaoqimenstorequeryRequest 初始化TaobaoqimenstorequeryAPIRequest对象
func NewTaobaoqimenstorequeryRequest() *TaobaoqimenstorequeryAPIRequest {
	return &TaobaoqimenstorequeryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoqimenstorequeryAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.store.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoqimenstorequeryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoqimenstorequeryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStoreId is StoreId Setter
// 已分配的线上门店ID
func (r *TaobaoqimenstorequeryAPIRequest) SetStoreId(_storeId int64) error {
	r._storeId = _storeId
	r.Set("store_id", _storeId)
	return nil
}

// GetStoreId StoreId Getter
func (r TaobaoqimenstorequeryAPIRequest) GetStoreId() int64 {
	return r._storeId
}
