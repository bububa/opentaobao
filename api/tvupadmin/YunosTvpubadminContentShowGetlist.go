package tvupadmin

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tvupadmin"
)

/* 
节目审核获取节目列表 
yunos.tvpubadmin.content.show.getlist

节目审核获取节目列表
*/
func YunosTvpubadminContentShowGetlist(clt *core.SDKClient, req *tvupadmin.YunosTvpubadminContentShowGetlistAPIRequest, session string) (*tvupadmin.YunosTvpubadminContentShowGetlistAPIResponse, error) {
    var resp tvupadmin.YunosTvpubadminContentShowGetlistAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
