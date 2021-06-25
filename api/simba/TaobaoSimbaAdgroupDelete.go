package simba

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/simba"
)

/* 
删除一个推广组 
taobao.simba.adgroup.delete

删除一个推广组
*/
func TaobaoSimbaAdgroupDelete(clt *core.SDKClient, req *simba.TaobaoSimbaAdgroupDeleteRequest, session string) (*simba.TaobaoSimbaAdgroupDeleteResponse, error) {
    var resp simba.TaobaoSimbaAdgroupDeleteAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
