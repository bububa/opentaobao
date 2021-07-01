package tmallcar

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tmallcar"
)

/* 
全量车型导入 
taobao.car.vehicleinfo.register

全量车型导入
*/
func TaobaoCarVehicleinfoRegister(clt *core.SDKClient, req *tmallcar.TaobaoCarVehicleinfoRegisterAPIRequest, session string) (*tmallcar.TaobaoCarVehicleinfoRegisterAPIResponse, error) {
    var resp tmallcar.TaobaoCarVehicleinfoRegisterAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
