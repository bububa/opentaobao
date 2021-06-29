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
    alipayUserId   string
    // 预约id
    orderId   string
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
func (r *AlibabaHealthVaccinNoticeOrderSignRequest) SetAlipayUserId(alipayUserId string) error {
    r.alipayUserId = alipayUserId
    r.Set("alipay_user_id", alipayUserId)
    return nil
}

// AlipayUserId Getter
func (r AlibabaHealthVaccinNoticeOrderSignRequest) GetAlipayUserId() string {
    return r.alipayUserId
}
// OrderId Setter
// 预约id
func (r *AlibabaHealthVaccinNoticeOrderSignRequest) SetOrderId(orderId string) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

// OrderId Getter
func (r AlibabaHealthVaccinNoticeOrderSignRequest) GetOrderId() string {
    return r.orderId
}
