package tmallcarenter

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallcarenter"
)

// Tmallcarcentervehicleinforegister 车型数据更新
// tmall.carcenter.vehicleinfo.register
//
// 基本车型信息维护
func Tmallcarcentervehicleinforegister(clt *core.SDKClient, req *tmallcarenter.TmallcarcentervehicleinforegisterAPIRequest, session string) (*tmallcarenter.TmallcarcentervehicleinforegisterAPIResponse, error) {
	var resp tmallcarenter.TmallcarcentervehicleinforegisterAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
