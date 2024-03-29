package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWdkEquipmentConveyorBatchconfirmAPIRequest 五道口悬挂链信息批量确认 API请求
// taobao.wdk.equipment.conveyor.batchconfirm
//
// 批量消息确认
type TaobaoWdkEquipmentConveyorBatchconfirmAPIRequest struct {
	model.Params
	// 待确认的uuid列表
	_uuids []string
	// 仓库code
	_warehouseCode string
}

// NewTaobaoWdkEquipmentConveyorBatchconfirmRequest 初始化TaobaoWdkEquipmentConveyorBatchconfirmAPIRequest对象
func NewTaobaoWdkEquipmentConveyorBatchconfirmRequest() *TaobaoWdkEquipmentConveyorBatchconfirmAPIRequest {
	return &TaobaoWdkEquipmentConveyorBatchconfirmAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoWdkEquipmentConveyorBatchconfirmAPIRequest) Reset() {
	r._uuids = r._uuids[:0]
	r._warehouseCode = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoWdkEquipmentConveyorBatchconfirmAPIRequest) GetApiMethodName() string {
	return "taobao.wdk.equipment.conveyor.batchconfirm"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoWdkEquipmentConveyorBatchconfirmAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoWdkEquipmentConveyorBatchconfirmAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUuids is Uuids Setter
// 待确认的uuid列表
func (r *TaobaoWdkEquipmentConveyorBatchconfirmAPIRequest) SetUuids(_uuids []string) error {
	r._uuids = _uuids
	r.Set("uuids", _uuids)
	return nil
}

// GetUuids Uuids Getter
func (r TaobaoWdkEquipmentConveyorBatchconfirmAPIRequest) GetUuids() []string {
	return r._uuids
}

// SetWarehouseCode is WarehouseCode Setter
// 仓库code
func (r *TaobaoWdkEquipmentConveyorBatchconfirmAPIRequest) SetWarehouseCode(_warehouseCode string) error {
	r._warehouseCode = _warehouseCode
	r.Set("warehouse_code", _warehouseCode)
	return nil
}

// GetWarehouseCode WarehouseCode Getter
func (r TaobaoWdkEquipmentConveyorBatchconfirmAPIRequest) GetWarehouseCode() string {
	return r._warehouseCode
}

var poolTaobaoWdkEquipmentConveyorBatchconfirmAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoWdkEquipmentConveyorBatchconfirmRequest()
	},
}

// GetTaobaoWdkEquipmentConveyorBatchconfirmRequest 从 sync.Pool 获取 TaobaoWdkEquipmentConveyorBatchconfirmAPIRequest
func GetTaobaoWdkEquipmentConveyorBatchconfirmAPIRequest() *TaobaoWdkEquipmentConveyorBatchconfirmAPIRequest {
	return poolTaobaoWdkEquipmentConveyorBatchconfirmAPIRequest.Get().(*TaobaoWdkEquipmentConveyorBatchconfirmAPIRequest)
}

// ReleaseTaobaoWdkEquipmentConveyorBatchconfirmAPIRequest 将 TaobaoWdkEquipmentConveyorBatchconfirmAPIRequest 放入 sync.Pool
func ReleaseTaobaoWdkEquipmentConveyorBatchconfirmAPIRequest(v *TaobaoWdkEquipmentConveyorBatchconfirmAPIRequest) {
	v.Reset()
	poolTaobaoWdkEquipmentConveyorBatchconfirmAPIRequest.Put(v)
}
