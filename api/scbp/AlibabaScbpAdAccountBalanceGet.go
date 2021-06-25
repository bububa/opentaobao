package scbp

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/scbp"
)

/* 
查询账户余额 
alibaba.scbp.ad.account.balance.get

查询推广账户余额
*/
func AlibabaScbpAdAccountBalanceGet(clt *core.SDKClient, req *scbp.AlibabaScbpAdAccountBalanceGetRequest, session string) (*scbp.AlibabaScbpAdAccountBalanceGetResponse, error) {
    var resp scbp.AlibabaScbpAdAccountBalanceGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
