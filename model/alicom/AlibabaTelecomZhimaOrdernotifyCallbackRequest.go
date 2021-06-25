package alicom

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
阿里通信芝麻订单通知 APIRequest
alibaba.telecom.zhima.ordernotify.callback

商家通知阿里通信，芝麻订单状态，阿里通信侧进行代扣支付、发货并确认收货
*/
type AlibabaTelecomZhimaOrdernotifyCallbackRequest struct {
    model.Params

    // 入参对象
    param0   *OrderStatusNotifyRequest 

}

func NewAlibabaTelecomZhimaOrdernotifyCallbackRequest() *AlibabaTelecomZhimaOrdernotifyCallbackRequest{
    return &AlibabaTelecomZhimaOrdernotifyCallbackRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaTelecomZhimaOrdernotifyCallbackRequest) GetApiMethodName() string {
    return "alibaba.telecom.zhima.ordernotify.callback"
}

func (r AlibabaTelecomZhimaOrdernotifyCallbackRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaTelecomZhimaOrdernotifyCallbackRequest) SetParam0(param0 *OrderStatusNotifyRequest) error {
    r.param0 = param0
    r.Set("param0", param0)
    return nil
}

func (r AlibabaTelecomZhimaOrdernotifyCallbackRequest) GetParam0() *OrderStatusNotifyRequest {
    return r.param0
}

