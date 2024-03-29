package fenxiao

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoRegionWarehouseQueryAPIRequest) Reset() {
	r._storeCode = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoRegionWarehouseQueryAPIRequest) GetApiMethodName() string {
	return "taobao.region.warehouse.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoRegionWarehouseQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoRegionWarehouseQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolTaobaoRegionWarehouseQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoRegionWarehouseQueryRequest()
	},
}

// GetTaobaoRegionWarehouseQueryRequest 从 sync.Pool 获取 TaobaoRegionWarehouseQueryAPIRequest
func GetTaobaoRegionWarehouseQueryAPIRequest() *TaobaoRegionWarehouseQueryAPIRequest {
	return poolTaobaoRegionWarehouseQueryAPIRequest.Get().(*TaobaoRegionWarehouseQueryAPIRequest)
}

// ReleaseTaobaoRegionWarehouseQueryAPIRequest 将 TaobaoRegionWarehouseQueryAPIRequest 放入 sync.Pool
func ReleaseTaobaoRegionWarehouseQueryAPIRequest(v *TaobaoRegionWarehouseQueryAPIRequest) {
	v.Reset()
	poolTaobaoRegionWarehouseQueryAPIRequest.Put(v)
}
