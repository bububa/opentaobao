package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoWdkEquipmentConveyorSystemeventGetAPIRequest
获取悬挂链系统事件 API请求
taobao.wdk.equipment.conveyor.systemevent.get

五道口悬挂链系统事件查询 */
type TaobaoWdkEquipmentConveyorSystemeventGetAPIRequest struct {
	model.Params
	// 仓库Id
	_warehouseId int64
	// 悬挂链Id，即wcsNum
	_conveyorId int64
	// 数据库id最小值
	_startId int64
}

// New
