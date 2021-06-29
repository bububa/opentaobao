package vaccin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
福州疫苗签到成功通知 APIRequest
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

func NewAlibabaHealthVaccinNoticeOrderSignRequest() *AlibabaHealthVaccinNoticeOrderSignRequest{
    return &AlibabaHealthVaccinNoticeOrderSignRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaHealthVaccinNoticeOrderSignRequest) GetApiMethodName() string {
    return "alibaba.health.vaccin.notice.order.sign"
}

func (r AlibabaHealthVaccinNoticeOrderSignRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaHealthVaccinNoticeOrderSignRequest) SetAlipayUserId(alipayUserId string) error {
    r.alipayUserId = alipayUserId
    r.Set("alipay_user_id", alipayUserId)
    return nil
}

func (r AlibabaHealthVaccinNoticeOrderSignRequest) GetAlipayUserId() string {
    return r.alipayUserId
}

func (r *AlibabaHealthVaccinNoticeOrderSignRequest) SetOrderId(orderId string) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

func (r AlibabaHealthVaccinNoticeOrderSignRequest) GetOrderId() string {
    return r.orderId
}

