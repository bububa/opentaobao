package tmallservice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
创建结算调整单 APIRequest
tmall.service.settleadjustment.request

提供给服务商在对结算有异议时，发起结算调整单。
通过说明工单ID，调整费用值，调整原因进行新建结算调整单。
*/
type TmallServiceSettleadjustmentRequestRequest struct {
    model.Params

    // 父节点
    paramSettleAdjustmentRequest   *SettleAdjustmentRequest 

}

func NewTmallServiceSettleadjustmentRequestRequest() *TmallServiceSettleadjustmentRequestRequest{
    return &TmallServiceSettleadjustmentRequestRequest{
        Params: model.NewParams(),
    }
}

func (r TmallServiceSettleadjustmentRequestRequest) GetApiMethodName() string {
    return "tmall.service.settleadjustment.request"
}

func (r TmallServiceSettleadjustmentRequestRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallServiceSettleadjustmentRequestRequest) SetParamSettleAdjustmentRequest(paramSettleAdjustmentRequest *SettleAdjustmentRequest) error {
    r.paramSettleAdjustmentRequest = paramSettleAdjustmentRequest
    r.Set("param_settle_adjustment_request", paramSettleAdjustmentRequest)
    return nil
}

func (r TmallServiceSettleadjustmentRequestRequest) GetParamSettleAdjustmentRequest() *SettleAdjustmentRequest {
    return r.paramSettleAdjustmentRequest
}

