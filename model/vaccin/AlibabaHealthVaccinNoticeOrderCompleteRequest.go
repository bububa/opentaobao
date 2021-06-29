package vaccin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
疫苗接种完成(带支付宝提醒) APIRequest
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

func NewAlibabaHealthVaccinNoticeOrderCompleteRequest() *AlibabaHealthVaccinNoticeOrderCompleteRequest{
    return &AlibabaHealthVaccinNoticeOrderCompleteRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaHealthVaccinNoticeOrderCompleteRequest) GetApiMethodName() string {
    return "alibaba.health.vaccin.notice.order.complete"
}

func (r AlibabaHealthVaccinNoticeOrderCompleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaHealthVaccinNoticeOrderCompleteRequest) SetAlipayUserId(alipayUserId string) error {
    r.alipayUserId = alipayUserId
    r.Set("alipay_user_id", alipayUserId)
    return nil
}

func (r AlibabaHealthVaccinNoticeOrderCompleteRequest) GetAlipayUserId() string {
    return r.alipayUserId
}

func (r *AlibabaHealthVaccinNoticeOrderCompleteRequest) SetOrderId(orderId string) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

func (r AlibabaHealthVaccinNoticeOrderCompleteRequest) GetOrderId() string {
    return r.orderId
}

