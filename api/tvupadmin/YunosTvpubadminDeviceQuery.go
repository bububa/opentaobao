package tvupadmin

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tvupadmin"
)

/* 
获取设备列表 
yunos.tvpubadmin.device.query

获取设备列表
*/
func YunosTvpubadminDeviceQuery(clt *core.SDKClient, req *tvupadmin.YunosTvpubadminDeviceQueryRequest, session string) (*tvupadmin.YunosTvpubadminDeviceQueryAPIResponse, error) {
    var resp tvupadmin.YunosTvpubadminDeviceQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
