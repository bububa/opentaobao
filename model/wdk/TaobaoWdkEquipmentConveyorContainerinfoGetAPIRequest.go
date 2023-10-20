package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaowdkequipmentconveyorcontainerinfogetAPIRequest 获取批次或波次中容器的信息 API请求
// taobao.wdk.equipment.conveyor.containerinfo.get
//
// 获取批次或波次中容器的信息
type TaobaowdkequipmentconveyorcontainerinfogetAPIRequest struct {
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

// NewTaobaowdkequipmentconveyorcontainerinfogetRequest 初始化TaobaowdkequipmentconveyorcontainerinfogetAPIRequest对象
func NewTaobaowdkequipmentconveyorcontainerinfogetRequest() *TaobaowdkequipmentconveyorcontainerinfogetAPIRequest {
	return &TaobaowdkequipmentconveyorcontainerinfogetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaowdkequipmentconveyorcontainerinfogetAPIRequest) GetApiMethodName() string {
	return "taobao.wdk.equipment.conveyor.containerinfo.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaowdkequipmentconveyorcontainerinfogetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaowdkequipmentconveyorcontainerinfogetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetContainerCode is ContainerCode Setter
// 容器号
func (r *TaobaowdkequipmentconveyorcontainerinfogetAPIRequest) SetContainerCode(_containerCode string) error {
	r._containerCode = _containerCode
	r.Set("container_code", _containerCode)
	return nil
}

// GetContainerCode ContainerCode Getter
func (r TaobaowdkequipmentconveyorcontainerinfogetAPIRequest) GetContainerCode() string {
	return r._containerCode
}

// SetBatchCode is BatchCode Setter
// 批次号，可以为空串
func (r *TaobaowdkequipmentconveyorcontainerinfogetAPIRequest) SetBatchCode(_batchCode string) error {
	r._batchCode = _batchCode
	r.Set("batch_code", _batchCode)
	return nil
}

// GetBatchCode BatchCode Getter
func (r TaobaowdkequipmentconveyorcontainerinfogetAPIRequest) GetBatchCode() string {
	return r._batchCode
}

// SetWaveCode is WaveCode Setter
// 波次号，可以为空串
func (r *TaobaowdkequipmentconveyorcontainerinfogetAPIRequest) SetWaveCode(_waveCode string) error {
	r._waveCode = _waveCode
	r.Set("wave_code", _waveCode)
	return nil
}

// GetWaveCode WaveCode Getter
func (r TaobaowdkequipmentconveyorcontainerinfogetAPIRequest) GetWaveCode() string {
	return r._waveCode
}

// SetWarehouseId is WarehouseId Setter
// 仓库id
func (r *TaobaowdkequipmentconveyorcontainerinfogetAPIRequest) SetWarehouseId(_warehouseId int64) error {
	r._warehouseId = _warehouseId
	r.Set("warehouse_id", _warehouseId)
	return nil
}

// GetWarehouseId WarehouseId Getter
func (r TaobaowdkequipmentconveyorcontainerinfogetAPIRequest) GetWarehouseId() int64 {
	return r._warehouseId
}
