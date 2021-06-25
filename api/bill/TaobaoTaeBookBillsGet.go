package bill

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/bill"
)

/* 
tae查询虚拟账户明细数据 
taobao.tae.book.bills.get

tae查询虚拟账户明细数据
*/
func TaobaoTaeBookBillsGet(clt *core.SDKClient, req *bill.TaobaoTaeBookBillsGetRequest, session string) (*bill.TaobaoTaeBookBillsGetResponse, error) {
    var resp bill.TaobaoTaeBookBillsGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
