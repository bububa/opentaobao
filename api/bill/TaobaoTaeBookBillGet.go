package bill

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/bill"
)

/* 
tae查询单笔虚拟账户明细 
taobao.tae.book.bill.get

tae查询单笔虚拟账户明细
*/
func TaobaoTaeBookBillGet(clt *core.SDKClient, req *bill.TaobaoTaeBookBillGetRequest, session string) (*bill.TaobaoTaeBookBillGetAPIResponse, error) {
    var resp bill.TaobaoTaeBookBillGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
