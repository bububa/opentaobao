package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWdkEquipmentConveyorSystemeventGetAPIRequest 获取悬挂链系统事件 API请求
// taobao.wdk.equipment.conveyor.systemevent.get
//
// 五道口悬挂链系统事件查询
type TaobaoWdkEquipmentConveyorSystemeventGetAPIRequest struct {
	model.Params
	// 仓库Id
	_warehouseId int64
	// 悬挂链Id，即wcsNum
	_conveyorId int64
	// 数据库id最小值
	_startId int64
}

// NewTaobaoWdkEquipmentConveyorSystemeventGetRequest 初始化TaobaoWdkEquipmentConveyorSystemeventGetAPIRequest对象
func NewTaobaoWdkEquipmentConveyorSystemeventGetRequest() *TaobaoWdkEquipmentConveyorSystemeventGetAPIRequest {
	return &TaobaoWdkEquipmentConveyorSystemeventGetAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoWdkEquipmentConveyorSystemeventGetAPIRequest) Reset() {
	r._warehouseId = 0
	r._conveyorId = 0
	r._startId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoWdkEquipmentConveyorSystemeventGetAPIRequest) GetApiMethodName() string {
	return "taobao.wdk.equipment.conveyor.systemevent.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoWdkEquipmentConveyorSystemeventGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoWdkEquipmentConveyorSystemeventGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWarehouseId is WarehouseId Setter
// 仓库Id
func (r *TaobaoWdkEquipmentConveyorSystemeventGetAPIRequest) SetWarehouseId(_warehouseId int64) error {
	r._warehouseId = _warehouseId
	r.Set("warehouse_id", _warehouseId)
	return nil
}

// GetWarehouseId WarehouseId Getter
func (r TaobaoWdkEquipmentConveyorSystemeventGetAPIRequest) GetWarehouseId() int64 {
	return r._warehouseId
}

// SetConveyorId is ConveyorId Setter
// 悬挂链Id，即wcsNum
func (r *TaobaoWdkEquipmentConveyorSystemeventGetAPIRequest) SetConveyorId(_conveyorId int64) error {
	r._conveyorId = _conveyorId
	r.Set("conveyor_id", _conveyorId)
	return nil
}

// GetConveyorId ConveyorId Getter
func (r TaobaoWdkEquipmentConveyorSystemeventGetAPIRequest) GetConveyorId() int64 {
	return r._conveyorId
}

// SetStartId is StartId Setter
// 数据库id最小值
func (r *TaobaoWdkEquipmentConveyorSystemeventGetAPIRequest) SetStartId(_startId int64) error {
	r._startId = _startId
	r.Set("start_id", _startId)
	return nil
}

// GetStartId StartId Getter
func (r TaobaoWdkEquipmentConveyorSystemeventGetAPIRequest) GetStartId() int64 {
	return r._startId
}

var poolTaobaoWdkEquipmentConveyorSystemeventGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoWdkEquipmentConveyorSystemeventGetRequest()
	},
}

// GetTaobaoWdkEquipmentConveyorSystemeventGetRequest 从 sync.Pool 获取 TaobaoWdkEquipmentConveyorSystemeventGetAPIRequest
func GetTaobaoWdkEquipmentConveyorSystemeventGetAPIRequest() *TaobaoWdkEquipmentConveyorSystemeventGetAPIRequest {
	return poolTaobaoWdkEquipmentConveyorSystemeventGetAPIRequest.Get().(*TaobaoWdkEquipmentConveyorSystemeventGetAPIRequest)
}

// ReleaseTaobaoWdkEquipmentConveyorSystemeventGetAPIRequest 将 TaobaoWdkEquipmentConveyorSystemeventGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoWdkEquipmentConveyorSystemeventGetAPIRequest(v *TaobaoWdkEquipmentConveyorSystemeventGetAPIRequest) {
	v.Reset()
	poolTaobaoWdkEquipmentConveyorSystemeventGetAPIRequest.Put(v)
}
