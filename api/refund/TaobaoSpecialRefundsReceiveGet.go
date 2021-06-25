package refund

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/refund"
)

/* 
特殊退款类型的纠纷单列表查询 
taobao.special.refunds.receive.get

特殊退款类型的纠纷单列表查询
*/
func TaobaoSpecialRefundsReceiveGet(clt *core.SDKClient, req *refund.TaobaoSpecialRefundsReceiveGetRequest, session string) (*refund.TaobaoSpecialRefundsReceiveGetResponse, error) {
    var resp refund.TaobaoSpecialRefundsReceiveGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
