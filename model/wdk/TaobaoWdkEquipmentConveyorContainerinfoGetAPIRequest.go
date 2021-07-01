package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoWdkEquipmentConveyorContainerinfoGetAPIRequest
获取批次或波次中容器的信息 API请求
taobao.wdk.equipment.conveyor.containerinfo.get

获取批次或波次中容器的信息 */
type TaobaoWdkEquipmentConveyorContainerinfoGetAPIRequest struct {
	model.Params
	// 仓库id
	_warehouseId int64
	// 容器号
	_containerCode string
	// 批次号，可以为空串
	_batchCode string
	// 波次号，可以为空串
	_waveCode string
}

// New
