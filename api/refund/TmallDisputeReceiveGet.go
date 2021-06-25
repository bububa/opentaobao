package refund

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/refund"
)

/* 
天猫逆向纠纷查询 
tmall.dispute.receive.get

展示商家所有退款信息
*/
func TmallDisputeReceiveGet(clt *core.SDKClient, req *refund.TmallDisputeReceiveGetRequest, session string) (*refund.TmallDisputeReceiveGetResponse, error) {
    var resp refund.TmallDisputeReceiveGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
