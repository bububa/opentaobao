package vaccin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
福州疫苗取消预约 API请求
alibaba.health.vaccin.notice.order.cancel

福州疫苗用户取消预约接口
*/
type AlibabaHealthVaccinNoticeOrderCancelRequest struct {
    model.Params
    // 支付宝用户id
    _alipayUserId   string
    // 预约id
    _orderId   string
}

// 初始化AlibabaHealthVaccinNoticeOrderCancelRequest对象
func NewAlibabaHealthVaccinNoticeOrderCancelRequest() *AlibabaHealthVaccinNoticeOrderCancelRequest{
    return &AlibabaHealthVaccinNoticeOrderCancelRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaHealthVaccinNoticeOrderCancelRequest) GetApiMethodName() string {
    return "alibaba.health.vaccin.notice.order.cancel"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaHealthVaccinNoticeOrderCancelRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AlipayUserId Setter
// 支付宝用户id
func (r *AlibabaHealthVaccinNoticeOrderCancelRequest) SetAlipayUserId(_alipayUserId string) error {
    r._alipayUserId = _alipayUserId
    r.Set("alipay_user_id", _alipayUserId)
    return nil
}

// AlipayUserId Getter
func (r AlibabaHealthVaccinNoticeOrderCancelRequest) GetAlipayUserId() string {
    return r._alipayUserId
}
// OrderId Setter
// 预约id
func (r *AlibabaHealthVaccinNoticeOrderCancelRequest) SetOrderId(_orderId string) error {
    r._orderId = _orderId
    r.Set("order_id", _orderId)
    return nil
}

// OrderId Getter
func (r AlibabaHealthVaccinNoticeOrderCancelRequest) GetOrderId() string {
    return r._orderId
}
