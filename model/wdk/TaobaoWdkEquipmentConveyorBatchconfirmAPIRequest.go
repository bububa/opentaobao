package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaowdkequipmentconveyorbatchconfirmAPIRequest 五道口悬挂链信息批量确认 API请求
// taobao.wdk.equipment.conveyor.batchconfirm
//
// 批量消息确认
type TaobaowdkequipmentconveyorbatchconfirmAPIRequest struct {
	model.Params
	// 待确认的uuid列表
	_uuids []string
	// 仓库code
	_warehouseCode string
}

// NewTaobaowdkequipmentconveyorbatchconfirmRequest 初始化TaobaowdkequipmentconveyorbatchconfirmAPIRequest对象
func NewTaobaowdkequipmentconveyorbatchconfirmRequest() *TaobaowdkequipmentconveyorbatchconfirmAPIRequest {
	return &TaobaowdkequipmentconveyorbatchconfirmAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaowdkequipmentconveyorbatchconfirmAPIRequest) GetApiMethodName() string {
	return "taobao.wdk.equipment.conveyor.batchconfirm"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaowdkequipmentconveyorbatchconfirmAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaowdkequipmentconveyorbatchconfirmAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUuids is Uuids Setter
// 待确认的uuid列表
func (r *TaobaowdkequipmentconveyorbatchconfirmAPIRequest) SetUuids(_uuids []string) error {
	r._uuids = _uuids
	r.Set("uuids", _uuids)
	return nil
}

// GetUuids Uuids Getter
func (r TaobaowdkequipmentconveyorbatchconfirmAPIRequest) GetUuids() []string {
	return r._uuids
}

// SetWarehouseCode is WarehouseCode Setter
// 仓库code
func (r *TaobaowdkequipmentconveyorbatchconfirmAPIRequest) SetWarehouseCode(_warehouseCode string) error {
	r._warehouseCode = _warehouseCode
	r.Set("warehouse_code", _warehouseCode)
	return nil
}

// GetWarehouseCode WarehouseCode Getter
func (r TaobaowdkequipmentconveyorbatchconfirmAPIRequest) GetWarehouseCode() string {
	return r._warehouseCode
}
