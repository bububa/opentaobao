package tvupadmin

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tvupadmin"
)

/* 
系统升级查询 
yunos.tvpubadmin.device.osupgradequery

系统升级查询
*/
func YunosTvpubadminDeviceOsupgradequery(clt *core.SDKClient, req *tvupadmin.YunosTvpubadminDeviceOsupgradequeryRequest, session string) (*tvupadmin.YunosTvpubadminDeviceOsupgradequeryAPIResponse, error) {
    var resp tvupadmin.YunosTvpubadminDeviceOsupgradequeryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
