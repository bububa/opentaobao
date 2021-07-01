package tvupadmin

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tvupadmin"
)

/* 
获取设备统计数据 
yunos.tvpubadmin.device.stats

获取设备统计数据
*/
func YunosTvpubadminDeviceStats(clt *core.SDKClient, req *tvupadmin.YunosTvpubadminDeviceStatsAPIRequest, session string) (*tvupadmin.YunosTvpubadminDeviceStatsAPIResponse, error) {
    var resp tvupadmin.YunosTvpubadminDeviceStatsAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
