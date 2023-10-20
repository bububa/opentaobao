package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWdkEquipmentConveyorConveyorinfoGetAPIRequest 获取五道口悬挂链信息 API请求
// taobao.wdk.equipment.conveyor.conveyorinfo.get
//
// 获取五道口悬挂链信息
type TaobaoWdkEquipmentConveyorConveyorinfoGetAPIRequest struct {
	model.Params
	// 仓库code
	_warehouseCode string
	// wcsNum
	_conveyorId int64
}

// NewTaobaoWdkEquipmentConveyorConveyorinfoGetRequest 初始化TaobaoWdkEquipmentConveyorConveyorinfoGetAPIRequest对象
func NewTaobaoWdkEquipmentConveyorConveyorinfoGetRequest() *TaobaoWdkEquipmentConveyorConveyorinfoGetAPIRequest {
	return &TaobaoWdkEquipmentConveyorConveyorinfoGetAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoWdkEquipmentConveyorConveyorinfoGetAPIRequest) Reset() {
	r._warehouseCode = ""
	r._conveyorId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoWdkEquipmentConveyorConveyorinfoGetAPIRequest) GetApiMethodName() string {
	return "taobao.wdk.equipment.conveyor.conveyorinfo.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoWdkEquipmentConveyorConveyorinfoGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoWdkEquipmentConveyorConveyorinfoGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWarehouseCode is WarehouseCode Setter
// 仓库code
func (r *TaobaoWdkEquipmentConveyorConveyorinfoGetAPIRequest) SetWarehouseCode(_warehouseCode string) error {
	r._warehouseCode = _warehouseCode
	r.Set("warehouse_code", _warehouseCode)
	return nil
}

// GetWarehouseCode WarehouseCode Getter
func (r TaobaoWdkEquipmentConveyorConveyorinfoGetAPIRequest) GetWarehouseCode() string {
	return r._warehouseCode
}

// SetConveyorId is ConveyorId Setter
// wcsNum
func (r *TaobaoWdkEquipmentConveyorConveyorinfoGetAPIRequest) SetConveyorId(_conveyorId int64) error {
	r._conveyorId = _conveyorId
	r.Set("conveyor_id", _conveyorId)
	return nil
}

// GetConveyorId ConveyorId Getter
func (r TaobaoWdkEquipmentConveyorConveyorinfoGetAPIRequest) GetConveyorId() int64 {
	return r._conveyorId
}

var poolTaobaoWdkEquipmentConveyorConveyorinfoGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoWdkEquipmentConveyorConveyorinfoGetRequest()
	},
}

// GetTaobaoWdkEquipmentConveyorConveyorinfoGetRequest 从 sync.Pool 获取 TaobaoWdkEquipmentConveyorConveyorinfoGetAPIRequest
func GetTaobaoWdkEquipmentConveyorConveyorinfoGetAPIRequest() *TaobaoWdkEquipmentConveyorConveyorinfoGetAPIRequest {
	return poolTaobaoWdkEquipmentConveyorConveyorinfoGetAPIRequest.Get().(*TaobaoWdkEquipmentConveyorConveyorinfoGetAPIRequest)
}

// ReleaseTaobaoWdkEquipmentConveyorConveyorinfoGetAPIRequest 将 TaobaoWdkEquipmentConveyorConveyorinfoGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoWdkEquipmentConveyorConveyorinfoGetAPIRequest(v *TaobaoWdkEquipmentConveyorConveyorinfoGetAPIRequest) {
	v.Reset()
	poolTaobaoWdkEquipmentConveyorConveyorinfoGetAPIRequest.Put(v)
}
