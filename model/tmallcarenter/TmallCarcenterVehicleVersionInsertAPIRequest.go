package tmallcarenter

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallCarcenterVehicleVersionInsertAPIRequest
汽车EPC版本压缩库新增接口 API请求
tmall.carcenter.vehicle.version.insert

汽车EPC版本压缩库新增接口 */
type TmallCarcenterVehicleVersionInsertAPIRequest struct {
	model.Params
	// 版本压缩库入参
	_dto *VersionVehicleInfoOriginalDto
}

// New
