package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoregionwarehousemanageAPIRequest 编辑仓库覆盖范围 API请求
// taobao.region.warehouse.manage
//
// 编辑仓库覆盖范围
type TaobaoregionwarehousemanageAPIRequest struct {
	model.Params
	// 可映射三级地址,例: 广东省
	_regions []string
	// 仓库编码
	_storeCode string
}

// NewTaobaoregionwarehousemanageRequest 初始化TaobaoregionwarehousemanageAPIRequest对象
func NewTaobaoregionwarehousemanageRequest() *TaobaoregionwarehousemanageAPIRequest {
	return &TaobaoregionwarehousemanageAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoregionwarehousemanageAPIRequest) GetApiMethodName() string {
	return "taobao.region.warehouse.manage"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoregionwarehousemanageAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoregionwarehousemanageAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRegions is Regions Setter
// 可映射三级地址,例: 广东省
func (r *TaobaoregionwarehousemanageAPIRequest) SetRegions(_regions []string) error {
	r._regions = _regions
	r.Set("regions", _regions)
	return nil
}

// GetRegions Regions Getter
func (r TaobaoregionwarehousemanageAPIRequest) GetRegions() []string {
	return r._regions
}

// SetStoreCode is StoreCode Setter
// 仓库编码
func (r *TaobaoregionwarehousemanageAPIRequest) SetStoreCode(_storeCode string) error {
	r._storeCode = _storeCode
	r.Set("store_code", _storeCode)
	return nil
}

// GetStoreCode StoreCode Getter
func (r TaobaoregionwarehousemanageAPIRequest) GetStoreCode() string {
	return r._storeCode
}
