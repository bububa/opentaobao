package ascp

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoLogisticsWarehouseCooperationBatchConfirmAPIRequest 仓合作关系确认 API请求
// taobao.logistics.warehouse.cooperation.batch.confirm
//
// 仓合作关系确认
type TaobaoLogisticsWarehouseCooperationBatchConfirmAPIRequest struct {
	model.Params
	// 仓合作关系确认入参
	_warehouseCooperationBatchConfirmRequest *WarehouseCooperationBatchConfirmRequest
}

// NewTaobaoLogisticsWarehouseCooperationBatchConfirmRequest 初始化TaobaoLogisticsWarehouseCooperationBatchConfirmAPIRequest对象
func NewTaobaoLogisticsWarehouseCooperationBatchConfirmRequest() *TaobaoLogisticsWarehouseCooperationBatchConfirmAPIRequest {
	return &TaobaoLogisticsWarehouseCooperationBatchConfirmAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoLogisticsWarehouseCooperationBatchConfirmAPIRequest) Reset() {
	r._warehouseCooperationBatchConfirmRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoLogisticsWarehouseCooperationBatchConfirmAPIRequest) GetApiMethodName() string {
	return "taobao.logistics.warehouse.cooperation.batch.confirm"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoLogisticsWarehouseCooperationBatchConfirmAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoLogisticsWarehouseCooperationBatchConfirmAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWarehouseCooperationBatchConfirmRequest is WarehouseCooperationBatchConfirmRequest Setter
// 仓合作关系确认入参
func (r *TaobaoLogisticsWarehouseCooperationBatchConfirmAPIRequest) SetWarehouseCooperationBatchConfirmRequest(_warehouseCooperationBatchConfirmRequest *WarehouseCooperationBatchConfirmRequest) error {
	r._warehouseCooperationBatchConfirmRequest = _warehouseCooperationBatchConfirmRequest
	r.Set("warehouse_cooperation_batch_confirm_request", _warehouseCooperationBatchConfirmRequest)
	return nil
}

// GetWarehouseCooperationBatchConfirmRequest WarehouseCooperationBatchConfirmRequest Getter
func (r TaobaoLogisticsWarehouseCooperationBatchConfirmAPIRequest) GetWarehouseCooperationBatchConfirmRequest() *WarehouseCooperationBatchConfirmRequest {
	return r._warehouseCooperationBatchConfirmRequest
}

var poolTaobaoLogisticsWarehouseCooperationBatchConfirmAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoLogisticsWarehouseCooperationBatchConfirmRequest()
	},
}

// GetTaobaoLogisticsWarehouseCooperationBatchConfirmRequest 从 sync.Pool 获取 TaobaoLogisticsWarehouseCooperationBatchConfirmAPIRequest
func GetTaobaoLogisticsWarehouseCooperationBatchConfirmAPIRequest() *TaobaoLogisticsWarehouseCooperationBatchConfirmAPIRequest {
	return poolTaobaoLogisticsWarehouseCooperationBatchConfirmAPIRequest.Get().(*TaobaoLogisticsWarehouseCooperationBatchConfirmAPIRequest)
}

// ReleaseTaobaoLogisticsWarehouseCooperationBatchConfirmAPIRequest 将 TaobaoLogisticsWarehouseCooperationBatchConfirmAPIRequest 放入 sync.Pool
func ReleaseTaobaoLogisticsWarehouseCooperationBatchConfirmAPIRequest(v *TaobaoLogisticsWarehouseCooperationBatchConfirmAPIRequest) {
	v.Reset()
	poolTaobaoLogisticsWarehouseCooperationBatchConfirmAPIRequest.Put(v)
}
