package trade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
外部渠道通知淘鲜达退款成功接口 API请求
alibaba.wdk.trade.refund.inform

该接口用于外部渠道退款成功后，通知淘鲜达底层履约完成退款流程。
*/
type AlibabaWdkTradeRefundInformRequest struct {
    model.Params
    // 通知退款成功请求
    informRefundSuccessRequest   *InformRefundSuccessRequest
}

// 初始化AlibabaWdkTradeRefundInformRequest对象
func NewAlibabaWdkTradeRefundInformRequest() *AlibabaWdkTradeRefundInformRequest{
    return &AlibabaWdkTradeRefundInformRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkTradeRefundInformRequest) GetApiMethodName() string {
    return "alibaba.wdk.trade.refund.inform"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkTradeRefundInformRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// InformRefundSuccessRequest Setter
// 通知退款成功请求
func (r *AlibabaWdkTradeRefundInformRequest) SetInformRefundSuccessRequest(informRefundSuccessRequest *InformRefundSuccessRequest) error {
    r.informRefundSuccessRequest = informRefundSuccessRequest
    r.Set("inform_refund_success_request", informRefundSuccessRequest)
    return nil
}

// InformRefundSuccessRequest Getter
func (r AlibabaWdkTradeRefundInformRequest) GetInformRefundSuccessRequest() *InformRefundSuccessRequest {
    return r.informRefundSuccessRequest
}
