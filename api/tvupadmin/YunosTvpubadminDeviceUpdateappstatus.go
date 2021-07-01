package tvupadmin

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tvupadmin"
)

/* 
更新应用版本审核状态 
yunos.tvpubadmin.device.updateappstatus

更新应用版本审核状态
*/
func YunosTvpubadminDeviceUpdateappstatus(clt *core.SDKClient, req *tvupadmin.YunosTvpubadminDeviceUpdateappstatusAPIRequest, session string) (*tvupadmin.YunosTvpubadminDeviceUpdateappstatusAPIResponse, error) {
    var resp tvupadmin.YunosTvpubadminDeviceUpdateappstatusAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
