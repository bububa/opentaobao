package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoinventorystorequeryAPIRequest 查询仓库信息 API请求
// taobao.inventory.store.query
//
// 查询商家仓信息
type TaobaoinventorystorequeryAPIRequest struct {
	model.Params
	// 商家的仓库编码
	_storeCode string
}

// NewTaobaoinventorystorequeryRequest 初始化TaobaoinventorystorequeryAPIRequest对象
func NewTaobaoinventorystorequeryRequest() *TaobaoinventorystorequeryAPIRequest {
	return &TaobaoinventorystorequeryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoinventorystorequeryAPIRequest) GetApiMethodName() string {
	return "taobao.inventory.store.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoinventorystorequeryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoinventorystorequeryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStoreCode is StoreCode Setter
// 商家的仓库编码
func (r *TaobaoinventorystorequeryAPIRequest) SetStoreCode(_storeCode string) error {
	r._storeCode = _storeCode
	r.Set("store_code", _storeCode)
	return nil
}

// GetStoreCode StoreCode Getter
func (r TaobaoinventorystorequeryAPIRequest) GetStoreCode() string {
	return r._storeCode
}
