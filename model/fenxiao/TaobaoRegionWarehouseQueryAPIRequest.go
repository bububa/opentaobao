package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoRegionWarehouseQueryAPIRequest 查询仓库覆盖范围 API请求
// taobao.region.warehouse.query
//
// 查询仓库覆盖范围
type TaobaoRegionWarehouseQueryAPIRequest struct {
	model.Params
	// 仓库编码
	_storeCode string
}

// NewTaobaoRegionWarehouseQueryRequest 初始化TaobaoRegionWarehouseQueryAPIRequest对象
func NewTaobaoRegionWarehouseQueryRequest() *TaobaoRegionWarehouseQueryAPIRequest {
	return &TaobaoRegionWarehouseQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoRegionWarehouseQueryAPIRequest) GetApiMethodName() string {
	return "taobao.region.warehouse.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoRegionWarehouseQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetStoreCode is StoreCode Setter
// 仓库编码
func (r *TaobaoRegionWarehouseQueryAPIRequest) SetStoreCode(_storeCode string) error {
	r._storeCode = _storeCode
	r.Set("store_code", _storeCode)
	return nil
}

// GetStoreCode StoreCode Getter
func (r TaobaoRegionWarehouseQueryAPIRequest) GetStoreCode() string {
	return r._storeCode
}
