package tmallservice

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tmallservice"
)

/* 
创建结算调整单 
tmall.service.settleadjustment.request

提供给服务商在对结算有异议时，发起结算调整单。
通过说明工单ID，调整费用值，调整原因进行新建结算调整单。
*/
func TmallServiceSettleadjustmentRequest(clt *core.SDKClient, req *tmallservice.TmallServiceSettleadjustmentRequestAPIRequest, session string) (*tmallservice.TmallServiceSettleadjustmentRequestAPIResponse, error) {
    var resp tmallservice.TmallServiceSettleadjustmentRequestAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
