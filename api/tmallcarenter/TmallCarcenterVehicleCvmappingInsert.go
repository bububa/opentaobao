package tmallcarenter

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tmallcarenter"
)

/* 
EPC车辆版本信息与底盘信息库关系绑定 
tmall.carcenter.vehicle.cvmapping.insert

EPC车辆版本信息与底盘信息库关系绑定
*/
func TmallCarcenterVehicleCvmappingInsert(clt *core.SDKClient, req *tmallcarenter.TmallCarcenterVehicleCvmappingInsertRequest, session string) (*tmallcarenter.TmallCarcenterVehicleCvmappingInsertAPIResponse, error) {
    var resp tmallcarenter.TmallCarcenterVehicleCvmappingInsertAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
