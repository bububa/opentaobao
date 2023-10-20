package tmallcarenter

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallcarenter"
)

// TmallCarcenterVehicleinfoRegister 车型数据更新
// tmall.carcenter.vehicleinfo.register
//
// 基本车型信息维护
func TmallCarcenterVehicleinfoRegister(clt *core.SDKClient, req *tmallcarenter.TmallCarcenterVehicleinfoRegisterAPIRequest, session string) (*tmallcarenter.TmallCarcenterVehicleinfoRegisterAPIResponse, error) {
	var resp tmallcarenter.TmallCarcenterVehicleinfoRegisterAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
