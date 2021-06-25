package bill

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/bill"
)

/* 
查询虚拟账户明细数据(自研发商家专用) 
taobao.bill.book.bills.get

查询虚拟账户明细数据
*/
func TaobaoBillBookBillsGet(clt *core.SDKClient, req *bill.TaobaoBillBookBillsGetRequest, session string) (*bill.TaobaoBillBookBillsGetResponse, error) {
    var resp bill.TaobaoBillBookBillsGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
