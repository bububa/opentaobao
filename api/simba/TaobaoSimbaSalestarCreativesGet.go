package simba

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/simba"
)

/* 
（新）批量获取创意 
taobao.simba.salestar.creatives.get

取得一个推广组的所有创意或者根据一个创意Id列表取得一组创意；<br/>如果同时提供了推广组Id和创意id列表，则优先使用推广组Id；
*/
func TaobaoSimbaSalestarCreativesGet(clt *core.SDKClient, req *simba.TaobaoSimbaSalestarCreativesGetRequest, session string) (*simba.TaobaoSimbaSalestarCreativesGetResponse, error) {
    var resp simba.TaobaoSimbaSalestarCreativesGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
