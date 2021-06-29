package tvupadmin

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tvupadmin"
)

/* 
查询终端信息 
yunos.tvpubadmin.device.tvid

通过UUID查询终端信息
*/
func YunosTvpubadminDeviceTvid(clt *core.SDKClient, req *tvupadmin.YunosTvpubadminDeviceTvidRequest, session string) (*tvupadmin.YunosTvpubadminDeviceTvidAPIResponse, error) {
    var resp tvupadmin.YunosTvpubadminDeviceTvidAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
