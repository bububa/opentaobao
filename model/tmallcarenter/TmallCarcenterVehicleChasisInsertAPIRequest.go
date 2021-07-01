package tmallcarenter

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallCarcenterVehicleChasisInsertAPIRequest
EPC车型底盘压缩库新增接口 API请求
tmall.carcenter.vehicle.chasis.insert

EPC车型底盘压缩库新增接口 */
type TmallCarcenterVehicleChasisInsertAPIRequest struct {
	model.Params
	// 底盘压缩库入参
	_dto *ChasisVehicleInfoOriginalDto
}

// New
