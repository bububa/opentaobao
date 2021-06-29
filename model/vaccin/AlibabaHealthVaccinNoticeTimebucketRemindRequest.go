package vaccin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
疫苗预约时间段提醒 APIRequest
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

func NewAlibabaHealthVaccinNoticeTimebucketRemindRequest() *AlibabaHealthVaccinNoticeTimebucketRemindRequest{
    return &AlibabaHealthVaccinNoticeTimebucketRemindRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaHealthVaccinNoticeTimebucketRemindRequest) GetApiMethodName() string {
    return "alibaba.health.vaccin.notice.timebucket.remind"
}

func (r AlibabaHealthVaccinNoticeTimebucketRemindRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaHealthVaccinNoticeTimebucketRemindRequest) SetAlipayUserId(alipayUserId string) error {
    r.alipayUserId = alipayUserId
    r.Set("alipay_user_id", alipayUserId)
    return nil
}

func (r AlibabaHealthVaccinNoticeTimebucketRemindRequest) GetAlipayUserId() string {
    return r.alipayUserId
}

func (r *AlibabaHealthVaccinNoticeTimebucketRemindRequest) SetVaccineName(vaccineName string) error {
    r.vaccineName = vaccineName
    r.Set("vaccine_name", vaccineName)
    return nil
}

func (r AlibabaHealthVaccinNoticeTimebucketRemindRequest) GetVaccineName() string {
    return r.vaccineName
}

func (r *AlibabaHealthVaccinNoticeTimebucketRemindRequest) SetReserveDate(reserveDate string) error {
    r.reserveDate = reserveDate
    r.Set("reserve_date", reserveDate)
    return nil
}

func (r AlibabaHealthVaccinNoticeTimebucketRemindRequest) GetReserveDate() string {
    return r.reserveDate
}

func (r *AlibabaHealthVaccinNoticeTimebucketRemindRequest) SetName(name string) error {
    r.name = name
    r.Set("name", name)
    return nil
}

func (r AlibabaHealthVaccinNoticeTimebucketRemindRequest) GetName() string {
    return r.name
}

func (r *AlibabaHealthVaccinNoticeTimebucketRemindRequest) SetTheTimes(theTimes string) error {
    r.theTimes = theTimes
    r.Set("the_times", theTimes)
    return nil
}

func (r AlibabaHealthVaccinNoticeTimebucketRemindRequest) GetTheTimes() string {
    return r.theTimes
}

func (r *AlibabaHealthVaccinNoticeTimebucketRemindRequest) SetPovStoreName(povStoreName string) error {
    r.povStoreName = povStoreName
    r.Set("pov_store_name", povStoreName)
    return nil
}

func (r AlibabaHealthVaccinNoticeTimebucketRemindRequest) GetPovStoreName() string {
    return r.povStoreName
}

func (r *AlibabaHealthVaccinNoticeTimebucketRemindRequest) SetReserveTime(reserveTime string) error {
    r.reserveTime = reserveTime
    r.Set("reserve_time", reserveTime)
    return nil
}

func (r AlibabaHealthVaccinNoticeTimebucketRemindRequest) GetReserveTime() string {
    return r.reserveTime
}

func (r *AlibabaHealthVaccinNoticeTimebucketRemindRequest) SetMobile(mobile string) error {
    r.mobile = mobile
    r.Set("mobile", mobile)
    return nil
}

func (r AlibabaHealthVaccinNoticeTimebucketRemindRequest) GetMobile() string {
    return r.mobile
}

