package bill

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/bill"
)

/* 
查询费用科目信息(限自研商家) 
taobao.bill.accounts.get

查询费用账户信息
*/
func TaobaoBillAccountsGet(clt *core.SDKClient, req *bill.TaobaoBillAccountsGetAPIRequest, session string) (*bill.TaobaoBillAccountsGetAPIResponse, error) {
    var resp bill.TaobaoBillAccountsGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
