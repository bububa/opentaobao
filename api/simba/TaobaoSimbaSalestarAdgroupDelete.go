package simba

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/simba"
)

/* 
(新)销量明星删除推广单元接口 
taobao.simba.salestar.adgroup.delete

删除一个推广组
*/
func TaobaoSimbaSalestarAdgroupDelete(clt *core.SDKClient, req *simba.TaobaoSimbaSalestarAdgroupDeleteAPIRequest, session string) (*simba.TaobaoSimbaSalestarAdgroupDeleteAPIResponse, error) {
    var resp simba.TaobaoSimbaSalestarAdgroupDeleteAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
