package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoWdkEquipmentConveyorExceptionslidewaylogGetAPIRequest
异常通道日志查询 API请求
taobao.wdk.equipment.conveyor.exceptionslidewaylog.get

五道口悬挂链异常通道事件查询 */
type TaobaoWdkEquipmentConveyorExceptionslidewaylogGetAPIRequest struct {
	model.Params
	// 仓库Id
	_warehouseId int64
	// 悬挂链Id，即wcsNum
	_conveyorId int64
	// 数据库id最小值
	_startId int64
}

// New
