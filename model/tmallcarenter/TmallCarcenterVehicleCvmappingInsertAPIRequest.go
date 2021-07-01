package tmallcarenter

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallCarcenterVehicleCvmappingInsertAPIRequest
EPC车辆版本信息与底盘信息库关系绑定 API请求
tmall.carcenter.vehicle.cvmapping.insert

EPC车辆版本信息与底盘信息库关系绑定 */
type TmallCarcenterVehicleCvmappingInsertAPIRequest struct {
	model.Params
	// 状态
	_status int64
	// 版本ID
	_supplierVersionCid string
	// 底盘ID
	_supplierChassisCid string
}

// New
