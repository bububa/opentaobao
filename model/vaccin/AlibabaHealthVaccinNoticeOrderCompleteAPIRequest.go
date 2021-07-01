package vaccin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
疫苗接种完成(带支付宝提醒) API请求
alibaba.health.vaccin.notice.order.complete

用户到店完成接种,ISV感知通知阿里健康完成接种,并通知用户!
*/
type AlibabaHealthVaccinNoticeOrderCompleteAPIRequest struct {
    model.Params
    // 支付宝唯一标识
    _alipayUserId   string
    // 在ISV预约单唯一标识
    _orderId   string
}

// 初始化AlibabaHealthVaccinNoticeOrderCompleteAPIRequest对象
func NewAlibabaHealthVaccinNoticeOrderCompleteRequest() *AlibabaHealthVaccinNoticeOrderCompleteAPIRequest{
    return &AlibabaHealthVaccinNoticeOrderCompleteAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaHealthVaccinNoticeOrderCompleteAPIRequest) GetApiMethodName() string {
    return "alibaba.health.vaccin.notice.order.complete"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaHealthVaccinNoticeOrderCompleteAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AlipayUserId Setter
// 支付宝唯一标识
func (r *AlibabaHealthVaccinNoticeOrderCompleteAPIRequest) SetAlipayUserId(_alipayUserId string) error {
    r._alipayUserId = _alipayUserId
    r.Set("alipay_user_id", _alipayUserId)
    return nil
}

// AlipayUserId Getter
func (r AlibabaHealthVaccinNoticeOrderCompleteAPIRequest) GetAlipayUserId() string {
    return r._alipayUserId
}
// OrderId Setter
// 在ISV预约单唯一标识
func (r *AlibabaHealthVaccinNoticeOrderCompleteAPIRequest) SetOrderId(_orderId string) error {
    r._orderId = _orderId
    r.Set("order_id", _orderId)
    return nil
}

// OrderId Getter
func (r AlibabaHealthVaccinNoticeOrderCompleteAPIRequest) GetOrderId() string {
    return r._orderId
}
