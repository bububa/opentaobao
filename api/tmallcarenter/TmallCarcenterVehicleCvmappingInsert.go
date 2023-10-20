package tmallcarenter

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallcarenter"
)

// Tmallcarcentervehiclecvmappinginsert EPC车辆版本信息与底盘信息库关系绑定
// tmall.carcenter.vehicle.cvmapping.insert
//
// EPC车辆版本信息与底盘信息库关系绑定
func Tmallcarcentervehiclecvmappinginsert(clt *core.SDKClient, req *tmallcarenter.TmallcarcentervehiclecvmappinginsertAPIRequest, session string) (*tmallcarenter.TmallcarcentervehiclecvmappinginsertAPIResponse, error) {
	var resp tmallcarenter.TmallcarcentervehiclecvmappinginsertAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
