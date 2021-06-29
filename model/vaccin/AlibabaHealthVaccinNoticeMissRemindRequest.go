package vaccin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
疫苗漏种提醒 API请求
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

// 初始化AlibabaHealthVaccinNoticeMissRemindRequest对象
func NewAlibabaHealthVaccinNoticeMissRemindRequest() *AlibabaHealthVaccinNoticeMissRemindRequest{
    return &AlibabaHealthVaccinNoticeMissRemindRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaHealthVaccinNoticeMissRemindRequest) GetApiMethodName() string {
    return "alibaba.health.vaccin.notice.miss.remind"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaHealthVaccinNoticeMissRemindRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AlipayUserId Setter
// 432421
func (r *AlibabaHealthVaccinNoticeMissRemindRequest) SetAlipayUserId(alipayUserId string) error {
    r.alipayUserId = alipayUserId
    r.Set("alipay_user_id", alipayUserId)
    return nil
}

// AlipayUserId Getter
func (r AlibabaHealthVaccinNoticeMissRemindRequest) GetAlipayUserId() string {
    return r.alipayUserId
}
// VaccineName Setter
// 多个疫苗英文逗号分隔
func (r *AlibabaHealthVaccinNoticeMissRemindRequest) SetVaccineName(vaccineName string) error {
    r.vaccineName = vaccineName
    r.Set("vaccine_name", vaccineName)
    return nil
}

// VaccineName Getter
func (r AlibabaHealthVaccinNoticeMissRemindRequest) GetVaccineName() string {
    return r.vaccineName
}
// ReserveDate Setter
// 2019-02-08 严格按照
func (r *AlibabaHealthVaccinNoticeMissRemindRequest) SetReserveDate(reserveDate string) error {
    r.reserveDate = reserveDate
    r.Set("reserve_date", reserveDate)
    return nil
}

// ReserveDate Getter
func (r AlibabaHealthVaccinNoticeMissRemindRequest) GetReserveDate() string {
    return r.reserveDate
}
// Name Setter
// 姓名
func (r *AlibabaHealthVaccinNoticeMissRemindRequest) SetName(name string) error {
    r.name = name
    r.Set("name", name)
    return nil
}

// Name Getter
func (r AlibabaHealthVaccinNoticeMissRemindRequest) GetName() string {
    return r.name
}
// Url Setter
// 点击提醒消息的跳转链接
func (r *AlibabaHealthVaccinNoticeMissRemindRequest) SetUrl(url string) error {
    r.url = url
    r.Set("url", url)
    return nil
}

// Url Getter
func (r AlibabaHealthVaccinNoticeMissRemindRequest) GetUrl() string {
    return r.url
}
