package tvupadmin

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tvupadmin"
)

/* 
获取停开服任务详情 
yunos.tvpubadmin.diccontroltask.getinfo

获取停开服任务详情
*/
func YunosTvpubadminDiccontroltaskGetinfo(clt *core.SDKClient, req *tvupadmin.YunosTvpubadminDiccontroltaskGetinfoRequest, session string) (*tvupadmin.YunosTvpubadminDiccontroltaskGetinfoAPIResponse, error) {
    var resp tvupadmin.YunosTvpubadminDiccontroltaskGetinfoAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
