package idle

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
闲鱼无忧购入仓模式服务商退款处理接口 APIRequest
alibaba.idle.isv.order.dealrefund

闲鱼无忧购业务入仓模式下，用户发起退款后，服务商使用此接口处理退款
*/
type AlibabaIdleIsvOrderDealrefundRequest struct {
    model.Params

    // 退款参数
    paramAppraiseIsvRefundRequest   *AppraiseIsvRefundRequest 

}

func NewAlibabaIdleIsvOrderDealrefundRequest() *AlibabaIdleIsvOrderDealrefundRequest{
    return &AlibabaIdleIsvOrderDealrefundRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaIdleIsvOrderDealrefundRequest) GetApiMethodName() string {
    return "alibaba.idle.isv.order.dealrefund"
}

func (r AlibabaIdleIsvOrderDealrefundRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaIdleIsvOrderDealrefundRequest) SetParamAppraiseIsvRefundRequest(paramAppraiseIsvRefundRequest *AppraiseIsvRefundRequest) error {
    r.paramAppraiseIsvRefundRequest = paramAppraiseIsvRefundRequest
    r.Set("param_appraise_isv_refund_request", paramAppraiseIsvRefundRequest)
    return nil
}

func (r AlibabaIdleIsvOrderDealrefundRequest) GetParamAppraiseIsvRefundRequest() *AppraiseIsvRefundRequest {
    return r.paramAppraiseIsvRefundRequest
}

