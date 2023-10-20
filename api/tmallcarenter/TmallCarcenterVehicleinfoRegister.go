package tmallcarenter

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallcarenter"
)

// TmallCarcenterVehicleinfoRegister 车型数据更新
// tmall.carcenter.vehicleinfo.register
//
// 基本车型信息维护
func TmallCarcenterVehicleinfoRegister(clt *core.SDKClient, req *tmallcarenter.TmallCarcenterVehicleinfoRegisterAPIRequest, resp *tmallcarenter.TmallCarcenterVehicleinfoRegisterAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
