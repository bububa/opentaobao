package simba

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/simba"
)

/* 
销量明星更新一个推广组的信息 
taobao.simba.salestar.adgroup.update

更新一个推广组的信息，可以设置 是否上线
*/
func TaobaoSimbaSalestarAdgroupUpdate(clt *core.SDKClient, req *simba.TaobaoSimbaSalestarAdgroupUpdateRequest, session string) (*simba.TaobaoSimbaSalestarAdgroupUpdateResponse, error) {
    var resp simba.TaobaoSimbaSalestarAdgroupUpdateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
