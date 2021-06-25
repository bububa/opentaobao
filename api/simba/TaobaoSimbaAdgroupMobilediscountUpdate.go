package simba

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/simba"
)

/* 
对推广组进行单独移动溢价 
taobao.simba.adgroup.mobilediscount.update

对推广组进行单独移动溢价
*/
func TaobaoSimbaAdgroupMobilediscountUpdate(clt *core.SDKClient, req *simba.TaobaoSimbaAdgroupMobilediscountUpdateRequest, session string) (*simba.TaobaoSimbaAdgroupMobilediscountUpdateResponse, error) {
    var resp simba.TaobaoSimbaAdgroupMobilediscountUpdateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
