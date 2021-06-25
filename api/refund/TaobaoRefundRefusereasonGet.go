package refund

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/refund"
)

/* 
获取拒绝原因列表 
taobao.refund.refusereason.get

获取商家拒绝原因列表
*/
func TaobaoRefundRefusereasonGet(clt *core.SDKClient, req *refund.TaobaoRefundRefusereasonGetRequest, session string) (*refund.TaobaoRefundRefusereasonGetResponse, error) {
    var resp refund.TaobaoRefundRefusereasonGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
