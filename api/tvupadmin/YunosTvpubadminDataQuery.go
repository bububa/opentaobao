package tvupadmin

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tvupadmin"
)

/* 
魔盒统计数据查询接口 
yunos.tvpubadmin.data.query

用于华数查询魔盒上的一些用户统计数据
*/
func YunosTvpubadminDataQuery(clt *core.SDKClient, req *tvupadmin.YunosTvpubadminDataQueryRequest, session string) (*tvupadmin.YunosTvpubadminDataQueryAPIResponse, error) {
    var resp tvupadmin.YunosTvpubadminDataQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}