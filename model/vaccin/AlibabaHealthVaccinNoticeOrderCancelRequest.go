package vaccin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
福州疫苗取消预约 APIRequest
alibaba.health.vaccin.notice.order.cancel

福州疫苗用户取消预约接口
*/
type AlibabaHealthVaccinNoticeOrderCancelRequest struct {
    model.Params

    // 支付宝用户id
    alipayUserId   string 

    // 预约id
    orderId   string 

}

func NewAlibabaHealthVaccinNoticeOrderCancelRequest() *AlibabaHealthVaccinNoticeOrderCancelRequest{
    return &AlibabaHealthVaccinNoticeOrderCancelRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaHealthVaccinNoticeOrderCancelRequest) GetApiMethodName() string {
    return "alibaba.health.vaccin.notice.order.cancel"
}

func (r AlibabaHealthVaccinNoticeOrderCancelRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaHealthVaccinNoticeOrderCancelRequest) SetAlipayUserId(alipayUserId string) error {
    r.alipayUserId = alipayUserId
    r.Set("alipay_user_id", alipayUserId)
    return nil
}

func (r AlibabaHealthVaccinNoticeOrderCancelRequest) GetAlipayUserId() string {
    return r.alipayUserId
}

func (r *AlibabaHealthVaccinNoticeOrderCancelRequest) SetOrderId(orderId string) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

func (r AlibabaHealthVaccinNoticeOrderCancelRequest) GetOrderId() string {
    return r.orderId
}

