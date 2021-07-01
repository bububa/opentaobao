package simba

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/simba"
)

/* 
批量删除adgroup的移动溢价 
taobao.simba.adgroup.mobilediscount.delete

批量删除adgroup的移动溢价
*/
func TaobaoSimbaAdgroupMobilediscountDelete(clt *core.SDKClient, req *simba.TaobaoSimbaAdgroupMobilediscountDeleteAPIRequest, session string) (*simba.TaobaoSimbaAdgroupMobilediscountDeleteAPIResponse, error) {
    var resp simba.TaobaoSimbaAdgroupMobilediscountDeleteAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
