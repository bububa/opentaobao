package tvupadmin

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tvupadmin"
)

/* 
获取停开服apk信息 
yunos.tvpubadmin.device.apkinfo

获取停开服apk信息
*/
func YunosTvpubadminDeviceApkinfo(clt *core.SDKClient, req *tvupadmin.YunosTvpubadminDeviceApkinfoAPIRequest, session string) (*tvupadmin.YunosTvpubadminDeviceApkinfoAPIResponse, error) {
    var resp tvupadmin.YunosTvpubadminDeviceApkinfoAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
