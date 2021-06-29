package tmallcarenter

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tmallcarenter"
)

/* 
EPC车型底盘压缩库新增接口 
tmall.carcenter.vehicle.chasis.insert

EPC车型底盘压缩库新增接口
*/
func TmallCarcenterVehicleChasisInsert(clt *core.SDKClient, req *tmallcarenter.TmallCarcenterVehicleChasisInsertRequest, session string) (*tmallcarenter.TmallCarcenterVehicleChasisInsertAPIResponse, error) {
    var resp tmallcarenter.TmallCarcenterVehicleChasisInsertAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
