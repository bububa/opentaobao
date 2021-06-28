package simba

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/simba"
)

/* 
创建一个推广组 
taobao.simba.adgroup.add

创建一个推广组
*/
func TaobaoSimbaAdgroupAdd(clt *core.SDKClient, req *simba.TaobaoSimbaAdgroupAddRequest, session string) (*simba.TaobaoSimbaAdgroupAddAPIResponse, error) {
    var resp simba.TaobaoSimbaAdgroupAddAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
