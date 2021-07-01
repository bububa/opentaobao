package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoWdkEquipmentConveyorStatusconfirmAPIRequest
悬挂链状态回传确认 API请求
taobao.wdk.equipment.conveyor.statusconfirm

悬挂链状态回传确认 */
type TaobaoWdkEquipmentConveyorStatusconfirmAPIRequest struct {
	model.Params
	// warehouse_code
	_warehouseCode string
	// uuid
	_uuid string
}

// New
