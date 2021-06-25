package simba

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/simba"
)

/* 
获取实时余额，”元”为单位 
taobao.simba.account.balance.get

获取实时余额，”元”为单位
*/
func TaobaoSimbaAccountBalanceGet(clt *core.SDKClient, req *simba.TaobaoSimbaAccountBalanceGetRequest, session string) (*simba.TaobaoSimbaAccountBalanceGetResponse, error) {
    var resp simba.TaobaoSimbaAccountBalanceGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
