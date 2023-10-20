package tmallcar

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallcar"
)

// TaobaoCarVehicleinfoRegister 全量车型导入
// taobao.car.vehicleinfo.register
//
// 全量车型导入
func TaobaoCarVehicleinfoRegister(clt *core.SDKClient, req *tmallcar.TaobaoCarVehicleinfoRegisterAPIRequest, resp *tmallcar.TaobaoCarVehicleinfoRegisterAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
