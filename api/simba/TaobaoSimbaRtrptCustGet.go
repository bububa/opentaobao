package simba

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/simba"
)

/* 
获取账户实时报表数据 
taobao.simba.rtrpt.cust.get

获取账户实时报表数据
*/
func TaobaoSimbaRtrptCustGet(clt *core.SDKClient, req *simba.TaobaoSimbaRtrptCustGetAPIRequest, session string) (*simba.TaobaoSimbaRtrptCustGetAPIResponse, error) {
    var resp simba.TaobaoSimbaRtrptCustGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
