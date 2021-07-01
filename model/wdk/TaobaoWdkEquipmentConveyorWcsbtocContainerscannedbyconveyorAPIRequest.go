package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoWdkEquipmentConveyorWcsbtocContainerscannedbyconveyorAPIRequest
容器被悬挂链扫描 API请求
taobao.wdk.equipment.conveyor.wcsbtoc.containerscannedbyconveyor

容器被悬挂链扫描 */
type TaobaoWdkEquipmentConveyorWcsbtocContainerscannedbyconveyorAPIRequest struct {
	model.Params
	// warehouse_code
	_warehouseCode string
	// wcs_num
	_wcsNum int64
}

// New
