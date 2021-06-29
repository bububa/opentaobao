package vaccin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
支付宝疫苗补种提醒信息 APIRequest
alibaba.health.vaccin.notice.replant.remind

支付宝疫苗补种提醒
*/
type AlibabaHealthVaccinNoticeReplantRemindRequest struct {
    model.Params

    // 支付宝ID
    alipayUserId   string 

    // 针次
    theTimes   string 

    // 预约id
    orderId   string 

}

func NewAlibabaHealthVaccinNoticeReplantRemindRequest() *AlibabaHealthVaccinNoticeReplantRemindRequest{
    return &AlibabaHealthVaccinNoticeReplantRemindRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaHealthVaccinNoticeReplantRemindRequest) GetApiMethodName() string {
    return "alibaba.health.vaccin.notice.replant.remind"
}

func (r AlibabaHealthVaccinNoticeReplantRemindRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaHealthVaccinNoticeReplantRemindRequest) SetAlipayUserId(alipayUserId string) error {
    r.alipayUserId = alipayUserId
    r.Set("alipay_user_id", alipayUserId)
    return nil
}

func (r AlibabaHealthVaccinNoticeReplantRemindRequest) GetAlipayUserId() string {
    return r.alipayUserId
}

func (r *AlibabaHealthVaccinNoticeReplantRemindRequest) SetTheTimes(theTimes string) error {
    r.theTimes = theTimes
    r.Set("the_times", theTimes)
    return nil
}

func (r AlibabaHealthVaccinNoticeReplantRemindRequest) GetTheTimes() string {
    return r.theTimes
}

func (r *AlibabaHealthVaccinNoticeReplantRemindRequest) SetOrderId(orderId string) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

func (r AlibabaHealthVaccinNoticeReplantRemindRequest) GetOrderId() string {
    return r.orderId
}

