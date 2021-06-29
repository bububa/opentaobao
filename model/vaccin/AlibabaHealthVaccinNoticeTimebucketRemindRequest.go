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
    _alipayUserId   string
    // 疫苗名称
    _vaccineName   string
    // 预约日期：2019-02-08 严格按照
    _reserveDate   string
    // 接种人姓名
    _name   string
    // 针次
    _theTimes   string
    // 接种点名称（通知方）
    _povStoreName   string
    // 可预约时段
    _reserveTime   string
    // 用户授权的手机号
    _mobile   string
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
func (r *AlibabaHealthVaccinNoticeTimebucketRemindRequest) SetAlipayUserId(_alipayUserId string) error {
    r._alipayUserId = _alipayUserId
    r.Set("alipay_user_id", _alipayUserId)
    return nil
}

// AlipayUserId Getter
func (r AlibabaHealthVaccinNoticeTimebucketRemindRequest) GetAlipayUserId() string {
    return r._alipayUserId
}
// VaccineName Setter
// 疫苗名称
func (r *AlibabaHealthVaccinNoticeTimebucketRemindRequest) SetVaccineName(_vaccineName string) error {
    r._vaccineName = _vaccineName
    r.Set("vaccine_name", _vaccineName)
    return nil
}

// VaccineName Getter
func (r AlibabaHealthVaccinNoticeTimebucketRemindRequest) GetVaccineName() string {
    return r._vaccineName
}
// ReserveDate Setter
// 预约日期：2019-02-08 严格按照
func (r *AlibabaHealthVaccinNoticeTimebucketRemindRequest) SetReserveDate(_reserveDate string) error {
    r._reserveDate = _reserveDate
    r.Set("reserve_date", _reserveDate)
    return nil
}

// ReserveDate Getter
func (r AlibabaHealthVaccinNoticeTimebucketRemindRequest) GetReserveDate() string {
    return r._reserveDate
}
// Name Setter
// 接种人姓名
func (r *AlibabaHealthVaccinNoticeTimebucketRemindRequest) SetName(_name string) error {
    r._name = _name
    r.Set("name", _name)
    return nil
}

// Name Getter
func (r AlibabaHealthVaccinNoticeTimebucketRemindRequest) GetName() string {
    return r._name
}
// TheTimes Setter
// 针次
func (r *AlibabaHealthVaccinNoticeTimebucketRemindRequest) SetTheTimes(_theTimes string) error {
    r._theTimes = _theTimes
    r.Set("the_times", _theTimes)
    return nil
}

// TheTimes Getter
func (r AlibabaHealthVaccinNoticeTimebucketRemindRequest) GetTheTimes() string {
    return r._theTimes
}
// PovStoreName Setter
// 接种点名称（通知方）
func (r *AlibabaHealthVaccinNoticeTimebucketRemindRequest) SetPovStoreName(_povStoreName string) error {
    r._povStoreName = _povStoreName
    r.Set("pov_store_name", _povStoreName)
    return nil
}

// PovStoreName Getter
func (r AlibabaHealthVaccinNoticeTimebucketRemindRequest) GetPovStoreName() string {
    return r._povStoreName
}
// ReserveTime Setter
// 可预约时段
func (r *AlibabaHealthVaccinNoticeTimebucketRemindRequest) SetReserveTime(_reserveTime string) error {
    r._reserveTime = _reserveTime
    r.Set("reserve_time", _reserveTime)
    return nil
}

// ReserveTime Getter
func (r AlibabaHealthVaccinNoticeTimebucketRemindRequest) GetReserveTime() string {
    return r._reserveTime
}
// Mobile Setter
// 用户授权的手机号
func (r *AlibabaHealthVaccinNoticeTimebucketRemindRequest) SetMobile(_mobile string) error {
    r._mobile = _mobile
    r.Set("mobile", _mobile)
    return nil
}

// Mobile Getter
func (r AlibabaHealthVaccinNoticeTimebucketRemindRequest) GetMobile() string {
    return r._mobile
}
