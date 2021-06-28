package bill

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/bill"
)

/* 
tae查询费用科目信息 
taobao.tae.accounts.get

tae查询费用科目信息
*/
func TaobaoTaeAccountsGet(clt *core.SDKClient, req *bill.TaobaoTaeAccountsGetRequest, session string) (*bill.TaobaoTaeAccountsGetAPIResponse, error) {
    var resp bill.TaobaoTaeAccountsGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
