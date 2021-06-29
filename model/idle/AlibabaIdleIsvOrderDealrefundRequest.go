package idle

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
闲鱼无忧购入仓模式服务商退款处理接口 API请求
alibaba.idle.isv.order.dealrefund

闲鱼无忧购业务入仓模式下，用户发起退款后，服务商使用此接口处理退款
*/
type AlibabaIdleIsvOrderDealrefundRequest struct {
    model.Params
    // 退款参数
    paramAppraiseIsvRefundRequest   *AppraiseIsvRefundRequest
}

// 初始化AlibabaIdleIsvOrderDealrefundRequest对象
func NewAlibabaIdleIsvOrderDealrefundRequest() *AlibabaIdleIsvOrderDealrefundRequest{
    return &AlibabaIdleIsvOrderDealrefundRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaIdleIsvOrderDealrefundRequest) GetApiMethodName() string {
    return "alibaba.idle.isv.order.dealrefund"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaIdleIsvOrderDealrefundRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamAppraiseIsvRefundRequest Setter
// 退款参数
func (r *AlibabaIdleIsvOrderDealrefundRequest) SetParamAppraiseIsvRefundRequest(paramAppraiseIsvRefundRequest *AppraiseIsvRefundRequest) error {
    r.paramAppraiseIsvRefundRequest = paramAppraiseIsvRefundRequest
    r.Set("param_appraise_isv_refund_request", paramAppraiseIsvRefundRequest)
    return nil
}

// ParamAppraiseIsvRefundRequest Getter
func (r AlibabaIdleIsvOrderDealrefundRequest) GetParamAppraiseIsvRefundRequest() *AppraiseIsvRefundRequest {
    return r.paramAppraiseIsvRefundRequest
}
