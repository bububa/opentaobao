package alicom

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
阿里通信芝麻订单通知 API请求
alibaba.telecom.zhima.ordernotify.callback

商家通知阿里通信，芝麻订单状态，阿里通信侧进行代扣支付、发货并确认收货
*/
type AlibabaTelecomZhimaOrdernotifyCallbackRequest struct {
    model.Params
    // 入参对象
    _param0   *OrderStatusNotifyRequest
}

// 初始化AlibabaTelecomZhimaOrdernotifyCallbackRequest对象
func NewAlibabaTelecomZhimaOrdernotifyCallbackRequest() *AlibabaTelecomZhimaOrdernotifyCallbackRequest{
    return &AlibabaTelecomZhimaOrdernotifyCallbackRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaTelecomZhimaOrdernotifyCallbackRequest) GetApiMethodName() string {
    return "alibaba.telecom.zhima.ordernotify.callback"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaTelecomZhimaOrdernotifyCallbackRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param0 Setter
// 入参对象
func (r *AlibabaTelecomZhimaOrdernotifyCallbackRequest) SetParam0(_param0 *OrderStatusNotifyRequest) error {
    r._param0 = _param0
    r.Set("param0", _param0)
    return nil
}

// Param0 Getter
func (r AlibabaTelecomZhimaOrdernotifyCallbackRequest) GetParam0() *OrderStatusNotifyRequest {
    return r._param0
}
