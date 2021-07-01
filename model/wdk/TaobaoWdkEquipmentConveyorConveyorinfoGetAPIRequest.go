package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoWdkEquipmentConveyorConveyorinfoGetAPIRequest
获取五道口悬挂链信息 API请求
taobao.wdk.equipment.conveyor.conveyorinfo.get

获取五道口悬挂链信息 */
type TaobaoWdkEquipmentConveyorConveyorinfoGetAPIRequest struct {
	model.Params
	// 仓库code
	_warehouseCode string
	// wcsNum
	_conveyorId int64
}

// New
