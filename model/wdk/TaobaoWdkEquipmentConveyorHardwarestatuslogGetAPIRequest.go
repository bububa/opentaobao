package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoWdkEquipmentConveyorHardwarestatuslogGetAPIRequest
硬件状态变化日志查询 API请求
taobao.wdk.equipment.conveyor.hardwarestatuslog.get

硬件状态变化日志查询 */
type TaobaoWdkEquipmentConveyorHardwarestatuslogGetAPIRequest struct {
	model.Params
	// 仓库Id
	_warehouseId int64
	// 悬挂链Id，即wcsNum
	_conveyorId int64
	// 数据库id最小值
	_startId int64
}

// New
