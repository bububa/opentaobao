package simba

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/simba"
)

/* 
获取推广词实时报表数据 
taobao.simba.rtrpt.bidword.get

获取推广词报表数据
*/
func TaobaoSimbaRtrptBidwordGet(clt *core.SDKClient, req *simba.TaobaoSimbaRtrptBidwordGetAPIRequest, session string) (*simba.TaobaoSimbaRtrptBidwordGetAPIResponse, error) {
    var resp simba.TaobaoSimbaRtrptBidwordGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
