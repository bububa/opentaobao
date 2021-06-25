package trade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
外部渠道通知淘鲜达退款成功接口 APIRequest
alibaba.wdk.trade.refund.inform

该接口用于外部渠道退款成功后，通知淘鲜达底层履约完成退款流程。
*/
type AlibabaWdkTradeRefundInformRequest struct {
    model.Params

    // 通知退款成功请求
    informRefundSuccessRequest   *InformRefundSuccessRequest 

}

func NewAlibabaWdkTradeRefundInformRequest() *AlibabaWdkTradeRefundInformRequest{
    return &AlibabaWdkTradeRefundInformRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkTradeRefundInformRequest) GetApiMethodName() string {
    return "alibaba.wdk.trade.refund.inform"
}

func (r AlibabaWdkTradeRefundInformRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkTradeRefundInformRequest) SetInformRefundSuccessRequest(informRefundSuccessRequest *InformRefundSuccessRequest) error {
    r.informRefundSuccessRequest = informRefundSuccessRequest
    r.Set("inform_refund_success_request", informRefundSuccessRequest)
    return nil
}

func (r AlibabaWdkTradeRefundInformRequest) GetInformRefundSuccessRequest() *InformRefundSuccessRequest {
    return r.informRefundSuccessRequest
}

