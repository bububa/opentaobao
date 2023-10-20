package tmallcarenter

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallcarenter"
)

// TmallCarcenterVehicleCvmappingInsert EPC车辆版本信息与底盘信息库关系绑定
// tmall.carcenter.vehicle.cvmapping.insert
//
// EPC车辆版本信息与底盘信息库关系绑定
func TmallCarcenterVehicleCvmappingInsert(clt *core.SDKClient, req *tmallcarenter.TmallCarcenterVehicleCvmappingInsertAPIRequest, resp *tmallcarenter.TmallCarcenterVehicleCvmappingInsertAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
