package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoregionwarehousequeryAPIRequest 查询仓库覆盖范围 API请求
// taobao.region.warehouse.query
//
// 查询仓库覆盖范围
type TaobaoregionwarehousequeryAPIRequest struct {
	model.Params
	// 仓库编码
	_storeCode string
}

// NewTaobaoregionwarehousequeryRequest 初始化TaobaoregionwarehousequeryAPIRequest对象
func NewTaobaoregionwarehousequeryRequest() *TaobaoregionwarehousequeryAPIRequest {
	return &TaobaoregionwarehousequeryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoregionwarehousequeryAPIRequest) GetApiMethodName() string {
	return "taobao.region.warehouse.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoregionwarehousequeryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoregionwarehousequeryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStoreCode is StoreCode Setter
// 仓库编码
func (r *TaobaoregionwarehousequeryAPIRequest) SetStoreCode(_storeCode string) error {
	r._storeCode = _storeCode
	r.Set("store_code", _storeCode)
	return nil
}

// GetStoreCode StoreCode Getter
func (r TaobaoregionwarehousequeryAPIRequest) GetStoreCode() string {
	return r._storeCode
}
