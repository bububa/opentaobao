package tvupadmin

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tvupadmin"
)

/* 
停开服任务列表 
yunos.tvpubadmin.diccontroltask.query

牌照方对终端设备的停开服管理
*/
func YunosTvpubadminDiccontroltaskQuery(clt *core.SDKClient, req *tvupadmin.YunosTvpubadminDiccontroltaskQueryAPIRequest, session string) (*tvupadmin.YunosTvpubadminDiccontroltaskQueryAPIResponse, error) {
    var resp tvupadmin.YunosTvpubadminDiccontroltaskQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
