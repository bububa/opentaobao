package simba

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/simba"
)

/* 
增加创意 
taobao.simba.creative.add

创建一个创意
*/
func TaobaoSimbaCreativeAdd(clt *core.SDKClient, req *simba.TaobaoSimbaCreativeAddAPIRequest, session string) (*simba.TaobaoSimbaCreativeAddAPIResponse, error) {
    var resp simba.TaobaoSimbaCreativeAddAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
