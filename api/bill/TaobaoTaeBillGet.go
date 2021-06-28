package bill

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/bill"
)

/* 
tae查询单笔账单明细 
taobao.tae.bill.get

查询单笔账单明细
*/
func TaobaoTaeBillGet(clt *core.SDKClient, req *bill.TaobaoTaeBillGetRequest, session string) (*bill.TaobaoTaeBillGetAPIResponse, error) {
    var resp bill.TaobaoTaeBillGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
