package vaccin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
疫苗漏种提醒 APIRequest
alibaba.health.vaccin.notice.miss.remind

医生消息提醒适龄儿童按计划接种
*/
type AlibabaHealthVaccinNoticeMissRemindRequest struct {
    model.Params

    // 432421
    alipayUserId   string 

    // 多个疫苗英文逗号分隔
    vaccineName   string 

    // 2019-02-08 严格按照
    reserveDate   string 

    // 姓名
    name   string 

    // 点击提醒消息的跳转链接
    url   string 

}

func NewAlibabaHealthVaccinNoticeMissRemindRequest() *AlibabaHealthVaccinNoticeMissRemindRequest{
    return &AlibabaHealthVaccinNoticeMissRemindRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaHealthVaccinNoticeMissRemindRequest) GetApiMethodName() string {
    return "alibaba.health.vaccin.notice.miss.remind"
}

func (r AlibabaHealthVaccinNoticeMissRemindRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaHealthVaccinNoticeMissRemindRequest) SetAlipayUserId(alipayUserId string) error {
    r.alipayUserId = alipayUserId
    r.Set("alipay_user_id", alipayUserId)
    return nil
}

func (r AlibabaHealthVaccinNoticeMissRemindRequest) GetAlipayUserId() string {
    return r.alipayUserId
}

func (r *AlibabaHealthVaccinNoticeMissRemindRequest) SetVaccineName(vaccineName string) error {
    r.vaccineName = vaccineName
    r.Set("vaccine_name", vaccineName)
    return nil
}

func (r AlibabaHealthVaccinNoticeMissRemindRequest) GetVaccineName() string {
    return r.vaccineName
}

func (r *AlibabaHealthVaccinNoticeMissRemindRequest) SetReserveDate(reserveDate string) error {
    r.reserveDate = reserveDate
    r.Set("reserve_date", reserveDate)
    return nil
}

func (r AlibabaHealthVaccinNoticeMissRemindRequest) GetReserveDate() string {
    return r.reserveDate
}

func (r *AlibabaHealthVaccinNoticeMissRemindRequest) SetName(name string) error {
    r.name = name
    r.Set("name", name)
    return nil
}

func (r AlibabaHealthVaccinNoticeMissRemindRequest) GetName() string {
    return r.name
}

func (r *AlibabaHealthVaccinNoticeMissRemindRequest) SetUrl(url string) error {
    r.url = url
    r.Set("url", url)
    return nil
}

func (r AlibabaHealthVaccinNoticeMissRemindRequest) GetUrl() string {
    return r.url
}

