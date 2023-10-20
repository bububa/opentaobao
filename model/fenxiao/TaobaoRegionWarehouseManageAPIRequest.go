package fenxiao

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoRegionWarehouseManageAPIRequest 编辑仓库覆盖范围 API请求
// taobao.region.warehouse.manage
//
// 编辑仓库覆盖范围
type TaobaoRegionWarehouseManageAPIRequest struct {
	model.Params
	// 可映射三级地址,例: 广东省
	_regions []string
	// 仓库编码
	_storeCode string
}

// NewTaobaoRegionWarehouseManageRequest 初始化TaobaoRegionWarehouseManageAPIRequest对象
func NewTaobaoRegionWarehouseManageRequest() *TaobaoRegionWarehouseManageAPIRequest {
	return &TaobaoRegionWarehouseManageAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoRegionWarehouseManageAPIRequest) Reset() {
	r._regions = r._regions[:0]
	r._storeCode = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoRegionWarehouseManageAPIRequest) GetApiMethodName() string {
	return "taobao.region.warehouse.manage"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoRegionWarehouseManageAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoRegionWarehouseManageAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRegions is Regions Setter
// 可映射三级地址,例: 广东省
func (r *TaobaoRegionWarehouseManageAPIRequest) SetRegions(_regions []string) error {
	r._regions = _regions
	r.Set("regions", _regions)
	return nil
}

// GetRegions Regions Getter
func (r TaobaoRegionWarehouseManageAPIRequest) GetRegions() []string {
	return r._regions
}

// SetStoreCode is StoreCode Setter
// 仓库编码
func (r *TaobaoRegionWarehouseManageAPIRequest) SetStoreCode(_storeCode string) error {
	r._storeCode = _storeCode
	r.Set("store_code", _storeCode)
	return nil
}

// GetStoreCode StoreCode Getter
func (r TaobaoRegionWarehouseManageAPIRequest) GetStoreCode() string {
	return r._storeCode
}

var poolTaobaoRegionWarehouseManageAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoRegionWarehouseManageRequest()
	},
}

// GetTaobaoRegionWarehouseManageRequest 从 sync.Pool 获取 TaobaoRegionWarehouseManageAPIRequest
func GetTaobaoRegionWarehouseManageAPIRequest() *TaobaoRegionWarehouseManageAPIRequest {
	return poolTaobaoRegionWarehouseManageAPIRequest.Get().(*TaobaoRegionWarehouseManageAPIRequest)
}

// ReleaseTaobaoRegionWarehouseManageAPIRequest 将 TaobaoRegionWarehouseManageAPIRequest 放入 sync.Pool
func ReleaseTaobaoRegionWarehouseManageAPIRequest(v *TaobaoRegionWarehouseManageAPIRequest) {
	v.Reset()
	poolTaobaoRegionWarehouseManageAPIRequest.Put(v)
}
