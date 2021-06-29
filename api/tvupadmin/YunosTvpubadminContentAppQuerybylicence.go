package tvupadmin

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tvupadmin"
)

/* 
按牌照查询应用 
yunos.tvpubadmin.content.app.querybylicence

按牌照查询应用
*/
func YunosTvpubadminContentAppQuerybylicence(clt *core.SDKClient, req *tvupadmin.YunosTvpubadminContentAppQuerybylicenceRequest, session string) (*tvupadmin.YunosTvpubadminContentAppQuerybylicenceAPIResponse, error) {
    var resp tvupadmin.YunosTvpubadminContentAppQuerybylicenceAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
