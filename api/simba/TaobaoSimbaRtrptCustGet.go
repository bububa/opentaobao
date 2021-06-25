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
func TaobaoSimbaRtrptCustGet(clt *core.SDKClient, req *simba.TaobaoSimbaRtrptCustGetRequest, session string) (*simba.TaobaoSimbaRtrptCustGetResponse, error) {
    var resp simba.TaobaoSimbaRtrptCustGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
