package refund

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/refund"
)

/* 
查询买家申请的退款列表 
taobao.refunds.apply.get

查询买家申请的退款列表，且查询外店的退款列表时需要指定交易类型
*/
func TaobaoRefundsApplyGet(clt *core.SDKClient, req *refund.TaobaoRefundsApplyGetRequest, session string) (*refund.TaobaoRefundsApplyGetAPIResponse, error) {
    var resp refund.TaobaoRefundsApplyGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
