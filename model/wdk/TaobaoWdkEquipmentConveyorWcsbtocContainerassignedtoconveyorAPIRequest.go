package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoWdkEquipmentConveyorWcsbtocContainerassignedtoconveyorAPIRequest
容器被预分拣器分配到悬挂链 API请求
taobao.wdk.equipment.conveyor.wcsbtoc.containerassignedtoconveyor

容器被预分拣器分配到悬挂链 */
type TaobaoWdkEquipmentConveyorWcsbtocContainerassignedtoconveyorAPIRequest struct {
	model.Params
	// warehouse_code
	_warehouseCode string
	// wcs_num
	_wcsNum int64
}

// New
