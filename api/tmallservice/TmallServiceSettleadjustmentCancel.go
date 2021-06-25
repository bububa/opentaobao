package tmallservice

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tmallservice"
)

/* 
取消结算调整单 
tmall.service.settleadjustment.cancel

提供给服务商在对取消已经发起的结算调整单。
通过说明调整单ID进行结算调整单取消。
*/
func TmallServiceSettleadjustmentCancel(clt *core.SDKClient, req *tmallservice.TmallServiceSettleadjustmentCancelRequest, session string) (*tmallservice.TmallServiceSettleadjustmentCancelResponse, error) {
    var resp tmallservice.TmallServiceSettleadjustmentCancelAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
