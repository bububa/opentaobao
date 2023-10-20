package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWdkEquipmentConveyorContainerinfoGetAPIRequest 获取批次或波次中容器的信息 API请求
// taobao.wdk.equipment.conveyor.containerinfo.get
//
// 获取批次或波次中容器的信息
type TaobaoWdkEquipmentConveyorContainerinfoGetAPIRequest struct {
	model.Params
	// 容器号
	_containerCode string
	// 批次号，可以为空串
	_batchCode string
	// 波次号，可以为空串
	_waveCode string
	// 仓库id
	_warehouseId int64
}

// NewTaobaoWdkEquipmentConveyorContainerinfoGetRequest 初始化TaobaoWdkEquipmentConveyorContainerinfoGetAPIRequest对象
func NewTaobaoWdkEquipmentConveyorContainerinfoGetRequest() *TaobaoWdkEquipmentConveyorContainerinfoGetAPIRequest {
	return &TaobaoWdkEquipmentConveyorContainerinfoGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoWdkEquipmentConveyorContainerinfoGetAPIRequest) GetApiMethodName() string {
	return "taobao.wdk.equipment.conveyor.containerinfo.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoWdkEquipmentConveyorContainerinfoGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoWdkEquipmentConveyorContainerinfoGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetContainerCode is ContainerCode Setter
// 容器号
func (r *TaobaoWdkEquipmentConveyorContainerinfoGetAPIRequest) SetContainerCode(_containerCode string) error {
	r._containerCode = _containerCode
	r.Set("container_code", _containerCode)
	return nil
}

// GetContainerCode ContainerCode Getter
func (r TaobaoWdkEquipmentConveyorContainerinfoGetAPIRequest) GetContainerCode() string {
	return r._containerCode
}

// SetBatchCode is BatchCode Setter
// 批次号，可以为空串
func (r *TaobaoWdkEquipmentConveyorContainerinfoGetAPIRequest) SetBatchCode(_batchCode string) error {
	r._batchCode = _batchCode
	r.Set("batch_code", _batchCode)
	return nil
}

// GetBatchCode BatchCode Getter
func (r TaobaoWdkEquipmentConveyorContainerinfoGetAPIRequest) GetBatchCode() string {
	return r._batchCode
}

// SetWaveCode is WaveCode Setter
// 波次号，可以为空串
func (r *TaobaoWdkEquipmentConveyorContainerinfoGetAPIRequest) SetWaveCode(_waveCode string) error {
	r._waveCode = _waveCode
	r.Set("wave_code", _waveCode)
	return nil
}

// GetWaveCode WaveCode Getter
func (r TaobaoWdkEquipmentConveyorContainerinfoGetAPIRequest) GetWaveCode() string {
	return r._waveCode
}

// SetWarehouseId is WarehouseId Setter
// 仓库id
func (r *TaobaoWdkEquipmentConveyorContainerinfoGetAPIRequest) SetWarehouseId(_warehouseId int64) error {
	r._warehouseId = _warehouseId
	r.Set("warehouse_id", _warehouseId)
	return nil
}

// GetWarehouseId WarehouseId Getter
func (r TaobaoWdkEquipmentConveyorContainerinfoGetAPIRequest) GetWarehouseId() int64 {
	return r._warehouseId
}
