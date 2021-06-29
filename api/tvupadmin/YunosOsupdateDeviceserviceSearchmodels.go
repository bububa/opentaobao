package tvupadmin

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tvupadmin"
)

/* 
根据关键词检索设备型号 
yunos.osupdate.deviceservice.searchmodels

根据关键词检索设备型号
*/
func YunosOsupdateDeviceserviceSearchmodels(clt *core.SDKClient, req *tvupadmin.YunosOsupdateDeviceserviceSearchmodelsRequest, session string) (*tvupadmin.YunosOsupdateDeviceserviceSearchmodelsAPIResponse, error) {
    var resp tvupadmin.YunosOsupdateDeviceserviceSearchmodelsAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
