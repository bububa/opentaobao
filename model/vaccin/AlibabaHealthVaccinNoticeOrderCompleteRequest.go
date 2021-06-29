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
type AlibabaHealthVaccinNoticeOrderCompleteRequest struct {
    model.Params
    // 支付宝唯一标识
    alipayUserId   string
    // 在ISV预约单唯一标识
    orderId   string
}

// 初始化AlibabaHealthVaccinNoticeOrderCompleteRequest对象
func NewAlibabaHealthVaccinNoticeOrderCompleteRequest() *AlibabaHealthVaccinNoticeOrderCompleteRequest{
    return &AlibabaHealthVaccinNoticeOrderCompleteRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaHealthVaccinNoticeOrderCompleteRequest) GetApiMethodName() string {
    return "alibaba.health.vaccin.notice.order.complete"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaHealthVaccinNoticeOrderCompleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AlipayUserId Setter
// 支付宝唯一标识
func (r *AlibabaHealthVaccinNoticeOrderCompleteRequest) SetAlipayUserId(alipayUserId string) error {
    r.alipayUserId = alipayUserId
    r.Set("alipay_user_id", alipayUserId)
    return nil
}

// AlipayUserId Getter
func (r AlibabaHealthVaccinNoticeOrderCompleteRequest) GetAlipayUserId() string {
    return r.alipayUserId
}
// OrderId Setter
// 在ISV预约单唯一标识
func (r *AlibabaHealthVaccinNoticeOrderCompleteRequest) SetOrderId(orderId string) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

// OrderId Getter
func (r AlibabaHealthVaccinNoticeOrderCompleteRequest) GetOrderId() string {
    return r.orderId
}
