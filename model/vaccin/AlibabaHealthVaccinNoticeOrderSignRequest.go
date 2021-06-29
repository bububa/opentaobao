package vaccin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
福州疫苗签到成功通知 API请求
alibaba.health.vaccin.notice.order.sign

福州疫苗用户签到成功记录
*/
type AlibabaHealthVaccinNoticeOrderSignRequest struct {
    model.Params
    // 支付宝用户id
    _alipayUserId   string
    // 预约id
    _orderId   string
}

// 初始化AlibabaHealthVaccinNoticeOrderSignRequest对象
func NewAlibabaHealthVaccinNoticeOrderSignRequest() *AlibabaHealthVaccinNoticeOrderSignRequest{
    return &AlibabaHealthVaccinNoticeOrderSignRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaHealthVaccinNoticeOrderSignRequest) GetApiMethodName() string {
    return "alibaba.health.vaccin.notice.order.sign"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaHealthVaccinNoticeOrderSignRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AlipayUserId Setter
// 支付宝用户id
func (r *AlibabaHealthVaccinNoticeOrderSignRequest) SetAlipayUserId(_alipayUserId string) error {
    r._alipayUserId = _alipayUserId
    r.Set("alipay_user_id", _alipayUserId)
    return nil
}

// AlipayUserId Getter
func (r AlibabaHealthVaccinNoticeOrderSignRequest) GetAlipayUserId() string {
    return r._alipayUserId
}
// OrderId Setter
// 预约id
func (r *AlibabaHealthVaccinNoticeOrderSignRequest) SetOrderId(_orderId string) error {
    r._orderId = _orderId
    r.Set("order_id", _orderId)
    return nil
}

// OrderId Getter
func (r AlibabaHealthVaccinNoticeOrderSignRequest) GetOrderId() string {
    return r._orderId
}
