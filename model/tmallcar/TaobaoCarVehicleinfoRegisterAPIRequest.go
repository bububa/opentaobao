package tmallcar

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoCarVehicleinfoRegisterAPIRequest
全量车型导入 API请求
taobao.car.vehicleinfo.register

全量车型导入 */
type TaobaoCarVehicleinfoRegisterAPIRequest struct {
	model.Params
	// 参数集合
	_paramList []FullInfoCarModelDto
}

// New
