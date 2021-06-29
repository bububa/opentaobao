package tmallservice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
创建结算调整单 API请求
tmall.service.settleadjustment.request

提供给服务商在对结算有异议时，发起结算调整单。
通过说明工单ID，调整费用值，调整原因进行新建结算调整单。
*/
type TmallServiceSettleadjustmentRequestRequest struct {
    model.Params
    // 父节点
    _paramSettleAdjustmentRequest   *SettleAdjustmentRequest
}

// 初始化TmallServiceSettleadjustmentRequestRequest对象
func NewTmallServiceSettleadjustmentRequestRequest() *TmallServiceSettleadjustmentRequestRequest{
    return &TmallServiceSettleadjustmentRequestRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallServiceSettleadjustmentRequestRequest) GetApiMethodName() string {
    return "tmall.service.settleadjustment.request"
}

// IRequest interface 方法, 获取API参数
func (r TmallServiceSettleadjustmentRequestRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamSettleAdjustmentRequest Setter
// 父节点
func (r *TmallServiceSettleadjustmentRequestRequest) SetParamSettleAdjustmentRequest(_paramSettleAdjustmentRequest *SettleAdjustmentRequest) error {
    r._paramSettleAdjustmentRequest = _paramSettleAdjustmentRequest
    r.Set("param_settle_adjustment_request", _paramSettleAdjustmentRequest)
    return nil
}

// ParamSettleAdjustmentRequest Getter
func (r TmallServiceSettleadjustmentRequestRequest) GetParamSettleAdjustmentRequest() *SettleAdjustmentRequest {
    return r._paramSettleAdjustmentRequest
}
