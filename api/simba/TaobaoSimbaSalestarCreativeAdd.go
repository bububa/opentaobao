package simba

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/simba"
)

/* 
（新）新建创意 
taobao.simba.salestar.creative.add

创建一个创意
*/
func TaobaoSimbaSalestarCreativeAdd(clt *core.SDKClient, req *simba.TaobaoSimbaSalestarCreativeAddAPIRequest, session string) (*simba.TaobaoSimbaSalestarCreativeAddAPIResponse, error) {
    var resp simba.TaobaoSimbaSalestarCreativeAddAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
