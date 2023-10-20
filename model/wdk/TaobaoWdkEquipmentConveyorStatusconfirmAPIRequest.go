package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWdkEquipmentConveyorStatusconfirmAPIRequest 悬挂链状态回传确认 API请求
// taobao.wdk.equipment.conveyor.statusconfirm
//
// 悬挂链状态回传确认
type TaobaoWdkEquipmentConveyorStatusconfirmAPIRequest struct {
	model.Params
	// warehouse_code
	_warehouseCode string
	// uuid
	_uuid string
}

// NewTaobaoWdkEquipmentConveyorStatusconfirmRequest 初始化TaobaoWdkEquipmentConveyorStatusconfirmAPIRequest对象
func NewTaobaoWdkEquipmentConveyorStatusconfirmRequest() *TaobaoWdkEquipmentConveyorStatusconfirmAPIRequest {
	return &TaobaoWdkEquipmentConveyorStatusconfirmAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoWdkEquipmentConveyorStatusconfirmAPIRequest) Reset() {
	r._warehouseCode = ""
	r._uuid = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoWdkEquipmentConveyorStatusconfirmAPIRequest) GetApiMethodName() string {
	return "taobao.wdk.equipment.conveyor.statusconfirm"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoWdkEquipmentConveyorStatusconfirmAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoWdkEquipmentConveyorStatusconfirmAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWarehouseCode is WarehouseCode Setter
// warehouse_code
func (r *TaobaoWdkEquipmentConveyorStatusconfirmAPIRequest) SetWarehouseCode(_warehouseCode string) error {
	r._warehouseCode = _warehouseCode
	r.Set("warehouse_code", _warehouseCode)
	return nil
}

// GetWarehouseCode WarehouseCode Getter
func (r TaobaoWdkEquipmentConveyorStatusconfirmAPIRequest) GetWarehouseCode() string {
	return r._warehouseCode
}

// SetUuid is Uuid Setter
// uuid
func (r *TaobaoWdkEquipmentConveyorStatusconfirmAPIRequest) SetUuid(_uuid string) error {
	r._uuid = _uuid
	r.Set("uuid", _uuid)
	return nil
}

// GetUuid Uuid Getter
func (r TaobaoWdkEquipmentConveyorStatusconfirmAPIRequest) GetUuid() string {
	return r._uuid
}

var poolTaobaoWdkEquipmentConveyorStatusconfirmAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoWdkEquipmentConveyorStatusconfirmRequest()
	},
}

// GetTaobaoWdkEquipmentConveyorStatusconfirmRequest 从 sync.Pool 获取 TaobaoWdkEquipmentConveyorStatusconfirmAPIRequest
func GetTaobaoWdkEquipmentConveyorStatusconfirmAPIRequest() *TaobaoWdkEquipmentConveyorStatusconfirmAPIRequest {
	return poolTaobaoWdkEquipmentConveyorStatusconfirmAPIRequest.Get().(*TaobaoWdkEquipmentConveyorStatusconfirmAPIRequest)
}

// ReleaseTaobaoWdkEquipmentConveyorStatusconfirmAPIRequest 将 TaobaoWdkEquipmentConveyorStatusconfirmAPIRequest 放入 sync.Pool
func ReleaseTaobaoWdkEquipmentConveyorStatusconfirmAPIRequest(v *TaobaoWdkEquipmentConveyorStatusconfirmAPIRequest) {
	v.Reset()
	poolTaobaoWdkEquipmentConveyorStatusconfirmAPIRequest.Put(v)
}
