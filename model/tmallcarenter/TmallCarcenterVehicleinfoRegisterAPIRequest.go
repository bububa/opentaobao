package tmallcarenter

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallCarcenterVehicleinfoRegisterAPIRequest
车型数据更新 API请求
tmall.carcenter.vehicleinfo.register

基本车型信息维护 */
type TmallCarcenterVehicleinfoRegisterAPIRequest struct {
	model.Params
	// 车型数据对象
	_vehicleInfo *OriginVehicleInfoDto
}

// New
