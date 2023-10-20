package ascp

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoLogisticsWarehouseCooperationUpdateAPIRequest 合作商家信息同步 API请求
// taobao.logistics.warehouse.cooperation.update
//
// 合作商家信息同步
type TaobaoLogisticsWarehouseCooperationUpdateAPIRequest struct {
	model.Params
	// 合作商家信息同步入参
	_warehouseCooperationUpdateRequest *WarehouseCooperationUpdateRequest
}

// NewTaobaoLogisticsWarehouseCooperationUpdateRequest 初始化TaobaoLogisticsWarehouseCooperationUpdateAPIRequest对象
func NewTaobaoLogisticsWarehouseCooperationUpdateRequest() *TaobaoLogisticsWarehouseCooperationUpdateAPIRequest {
	return &TaobaoLogisticsWarehouseCooperationUpdateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoLogisticsWarehouseCooperationUpdateAPIRequest) Reset() {
	r._warehouseCooperationUpdateRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoLogisticsWarehouseCooperationUpdateAPIRequest) GetApiMethodName() string {
	return "taobao.logistics.warehouse.cooperation.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoLogisticsWarehouseCooperationUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoLogisticsWarehouseCooperationUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWarehouseCooperationUpdateRequest is WarehouseCooperationUpdateRequest Setter
// 合作商家信息同步入参
func (r *TaobaoLogisticsWarehouseCooperationUpdateAPIRequest) SetWarehouseCooperationUpdateRequest(_warehouseCooperationUpdateRequest *WarehouseCooperationUpdateRequest) error {
	r._warehouseCooperationUpdateRequest = _warehouseCooperationUpdateRequest
	r.Set("warehouse_cooperation_update_request", _warehouseCooperationUpdateRequest)
	return nil
}

// GetWarehouseCooperationUpdateRequest WarehouseCooperationUpdateRequest Getter
func (r TaobaoLogisticsWarehouseCooperationUpdateAPIRequest) GetWarehouseCooperationUpdateRequest() *WarehouseCooperationUpdateRequest {
	return r._warehouseCooperationUpdateRequest
}

var poolTaobaoLogisticsWarehouseCooperationUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoLogisticsWarehouseCooperationUpdateRequest()
	},
}

// GetTaobaoLogisticsWarehouseCooperationUpdateRequest 从 sync.Pool 获取 TaobaoLogisticsWarehouseCooperationUpdateAPIRequest
func GetTaobaoLogisticsWarehouseCooperationUpdateAPIRequest() *TaobaoLogisticsWarehouseCooperationUpdateAPIRequest {
	return poolTaobaoLogisticsWarehouseCooperationUpdateAPIRequest.Get().(*TaobaoLogisticsWarehouseCooperationUpdateAPIRequest)
}

// ReleaseTaobaoLogisticsWarehouseCooperationUpdateAPIRequest 将 TaobaoLogisticsWarehouseCooperationUpdateAPIRequest 放入 sync.Pool
func ReleaseTaobaoLogisticsWarehouseCooperationUpdateAPIRequest(v *TaobaoLogisticsWarehouseCooperationUpdateAPIRequest) {
	v.Reset()
	poolTaobaoLogisticsWarehouseCooperationUpdateAPIRequest.Put(v)
}
