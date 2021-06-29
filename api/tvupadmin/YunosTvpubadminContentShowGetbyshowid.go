package tvupadmin

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tvupadmin"
)

/* 
迎客松根据节目id获取节目元数据 
yunos.tvpubadmin.content.show.getbyshowid

迎客松根据节目id获取节目元数据
*/
func YunosTvpubadminContentShowGetbyshowid(clt *core.SDKClient, req *tvupadmin.YunosTvpubadminContentShowGetbyshowidRequest, session string) (*tvupadmin.YunosTvpubadminContentShowGetbyshowidAPIResponse, error) {
    var resp tvupadmin.YunosTvpubadminContentShowGetbyshowidAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
