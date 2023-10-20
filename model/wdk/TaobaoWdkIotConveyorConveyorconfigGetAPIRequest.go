package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWdkIotConveyorConveyorconfigGetAPIRequest 获取悬挂链基本配置信息 API请求
// taobao.wdk.iot.conveyor.conveyorconfig.get
//
// 用于从云端WCS获取悬挂链基本配置信息
type TaobaoWdkIotConveyorConveyorconfigGetAPIRequest struct {
	model.Params
	// 仓编码
	_warehouseCode string
	// 悬挂链id，默认为1
	_conveyorId int64
}

// NewTaobaoWdkIotConveyorConveyorconfigGetRequest 初始化TaobaoWdkIotConveyorConveyorconfigGetAPIRequest对象
func NewTaobaoWdkIotConveyorConveyorconfigGetRequest() *TaobaoWdkIotConveyorConveyorconfigGetAPIRequest {
	return &TaobaoWdkIotConveyorConveyorconfigGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoWdkIotConveyorConveyorconfigGetAPIRequest) GetApiMethodName() string {
	return "taobao.wdk.iot.conveyor.conveyorconfig.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoWdkIotConveyorConveyorconfigGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoWdkIotConveyorConveyorconfigGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWarehouseCode is WarehouseCode Setter
// 仓编码
func (r *TaobaoWdkIotConveyorConveyorconfigGetAPIRequest) SetWarehouseCode(_warehouseCode string) error {
	r._warehouseCode = _warehouseCode
	r.Set("warehouse_code", _warehouseCode)
	return nil
}

// GetWarehouseCode WarehouseCode Getter
func (r TaobaoWdkIotConveyorConveyorconfigGetAPIRequest) GetWarehouseCode() string {
	return r._warehouseCode
}

// SetConveyorId is ConveyorId Setter
// 悬挂链id，默认为1
func (r *TaobaoWdkIotConveyorConveyorconfigGetAPIRequest) SetConveyorId(_conveyorId int64) error {
	r._conveyorId = _conveyorId
	r.Set("conveyor_id", _conveyorId)
	return nil
}

// GetConveyorId ConveyorId Getter
func (r TaobaoWdkIotConveyorConveyorconfigGetAPIRequest) GetConveyorId() int64 {
	return r._conveyorId
}
