package vaccin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
疫苗预约时间段提醒 API请求
alibaba.health.vaccin.notice.timebucket.remind

疫苗预约时间段提醒
*/
type AlibabaHealthVaccinNoticeTimebucketRemindRequest struct {
    model.Params
    // 432421
    alipayUserId   string
    // 疫苗名称
    vaccineName   string
    // 预约日期：2019-02-08 严格按照
    reserveDate   string
    // 接种人姓名
    name   string
    // 针次
    theTimes   string
    // 接种点名称（通知方）
    povStoreName   string
    // 可预约时段
    reserveTime   string
    // 用户授权的手机号
    mobile   string
}

// 初始化AlibabaHealthVaccinNoticeTimebucketRemindRequest对象
func NewAlibabaHealthVaccinNoticeTimebucketRemindRequest() *AlibabaHealthVaccinNoticeTimebucketRemindRequest{
    return &AlibabaHealthVaccinNoticeTimebucketRemindRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaHealthVaccinNoticeTimebucketRemindRequest) GetApiMethodName() string {
    return "alibaba.health.vaccin.notice.timebucket.remind"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaHealthVaccinNoticeTimebucketRemindRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AlipayUserId Setter
// 432421
func (r *AlibabaHealthVaccinNoticeTimebucketRemindRequest) SetAlipayUserId(alipayUserId string) error {
    r.alipayUserId = alipayUserId
    r.Set("alipay_user_id", alipayUserId)
    return nil
}

// AlipayUserId Getter
func (r AlibabaHealthVaccinNoticeTimebucketRemindRequest) GetAlipayUserId() string {
    return r.alipayUserId
}
// VaccineName Setter
// 疫苗名称
func (r *AlibabaHealthVaccinNoticeTimebucketRemindRequest) SetVaccineName(vaccineName string) error {
    r.vaccineName = vaccineName
    r.Set("vaccine_name", vaccineName)
    return nil
}

// VaccineName Getter
func (r AlibabaHealthVaccinNoticeTimebucketRemindRequest) GetVaccineName() string {
    return r.vaccineName
}
// ReserveDate Setter
// 预约日期：2019-02-08 严格按照
func (r *AlibabaHealthVaccinNoticeTimebucketRemindRequest) SetReserveDate(reserveDate string) error {
    r.reserveDate = reserveDate
    r.Set("reserve_date", reserveDate)
    return nil
}

// ReserveDate Getter
func (r AlibabaHealthVaccinNoticeTimebucketRemindRequest) GetReserveDate() string {
    return r.reserveDate
}
// Name Setter
// 接种人姓名
func (r *AlibabaHealthVaccinNoticeTimebucketRemindRequest) SetName(name string) error {
    r.name = name
    r.Set("name", name)
    return nil
}

// Name Getter
func (r AlibabaHealthVaccinNoticeTimebucketRemindRequest) GetName() string {
    return r.name
}
// TheTimes Setter
// 针次
func (r *AlibabaHealthVaccinNoticeTimebucketRemindRequest) SetTheTimes(theTimes string) error {
    r.theTimes = theTimes
    r.Set("the_times", theTimes)
    return nil
}

// TheTimes Getter
func (r AlibabaHealthVaccinNoticeTimebucketRemindRequest) GetTheTimes() string {
    return r.theTimes
}
// PovStoreName Setter
// 接种点名称（通知方）
func (r *AlibabaHealthVaccinNoticeTimebucketRemindRequest) SetPovStoreName(povStoreName string) error {
    r.povStoreName = povStoreName
    r.Set("pov_store_name", povStoreName)
    return nil
}

// PovStoreName Getter
func (r AlibabaHealthVaccinNoticeTimebucketRemindRequest) GetPovStoreName() string {
    return r.povStoreName
}
// ReserveTime Setter
// 可预约时段
func (r *AlibabaHealthVaccinNoticeTimebucketRemindRequest) SetReserveTime(reserveTime string) error {
    r.reserveTime = reserveTime
    r.Set("reserve_time", reserveTime)
    return nil
}

// ReserveTime Getter
func (r AlibabaHealthVaccinNoticeTimebucketRemindRequest) GetReserveTime() string {
    return r.reserveTime
}
// Mobile Setter
// 用户授权的手机号
func (r *AlibabaHealthVaccinNoticeTimebucketRemindRequest) SetMobile(mobile string) error {
    r.mobile = mobile
    r.Set("mobile", mobile)
    return nil
}

// Mobile Getter
func (r AlibabaHealthVaccinNoticeTimebucketRemindRequest) GetMobile() string {
    return r.mobile
}
