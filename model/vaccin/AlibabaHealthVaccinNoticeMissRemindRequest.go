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
    _alipayUserId   string
    // 多个疫苗英文逗号分隔
    _vaccineName   string
    // 2019-02-08 严格按照
    _reserveDate   string
    // 姓名
    _name   string
    // 点击提醒消息的跳转链接
    _url   string
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
func (r *AlibabaHealthVaccinNoticeMissRemindRequest) SetAlipayUserId(_alipayUserId string) error {
    r._alipayUserId = _alipayUserId
    r.Set("alipay_user_id", _alipayUserId)
    return nil
}

// AlipayUserId Getter
func (r AlibabaHealthVaccinNoticeMissRemindRequest) GetAlipayUserId() string {
    return r._alipayUserId
}
// VaccineName Setter
// 多个疫苗英文逗号分隔
func (r *AlibabaHealthVaccinNoticeMissRemindRequest) SetVaccineName(_vaccineName string) error {
    r._vaccineName = _vaccineName
    r.Set("vaccine_name", _vaccineName)
    return nil
}

// VaccineName Getter
func (r AlibabaHealthVaccinNoticeMissRemindRequest) GetVaccineName() string {
    return r._vaccineName
}
// ReserveDate Setter
// 2019-02-08 严格按照
func (r *AlibabaHealthVaccinNoticeMissRemindRequest) SetReserveDate(_reserveDate string) error {
    r._reserveDate = _reserveDate
    r.Set("reserve_date", _reserveDate)
    return nil
}

// ReserveDate Getter
func (r AlibabaHealthVaccinNoticeMissRemindRequest) GetReserveDate() string {
    return r._reserveDate
}
// Name Setter
// 姓名
func (r *AlibabaHealthVaccinNoticeMissRemindRequest) SetName(_name string) error {
    r._name = _name
    r.Set("name", _name)
    return nil
}

// Name Getter
func (r AlibabaHealthVaccinNoticeMissRemindRequest) GetName() string {
    return r._name
}
// Url Setter
// 点击提醒消息的跳转链接
func (r *AlibabaHealthVaccinNoticeMissRemindRequest) SetUrl(_url string) error {
    r._url = _url
    r.Set("url", _url)
    return nil
}

// Url Getter
func (r AlibabaHealthVaccinNoticeMissRemindRequest) GetUrl() string {
    return r._url
}
