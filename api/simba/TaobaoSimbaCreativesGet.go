package simba

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/simba"
)

/* 
批量获得创意 
taobao.simba.creatives.get

取得一个推广组的所有创意或者根据一个创意Id列表取得一组创意；<br/>如果同时提供了推广组Id和创意id列表，则优先使用推广组Id；
*/
func TaobaoSimbaCreativesGet(clt *core.SDKClient, req *simba.TaobaoSimbaCreativesGetRequest, session string) (*simba.TaobaoSimbaCreativesGetAPIResponse, error) {
    var resp simba.TaobaoSimbaCreativesGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
