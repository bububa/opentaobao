package vaccin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaHealthVaccinNoticeTimebucketRemindAPIRequest
疫苗预约时间段提醒 API请求
alibaba.health.vaccin.notice.timebucket.remind

疫苗预约时间段提醒 */
type AlibabaHealthVaccinNoticeTimebucketRemindAPIRequest struct {
	model.Params
	// 432421
	_alipayUserId string
	// 疫苗名称
	_vaccineName string
	// 预约日期：2019-02-08 严格按照
	_reserveDate string
	// 接种人姓名
	_name string
	// 针次
	_theTimes string
	// 接种点名称（通知方）
	_povStoreName string
	// 可预约时段
	_reserveTime string
	// 用户授权的手机号
	_mobile string
}

// NewAlibabaHealthVaccinNoticeTimebucketRemindRequest 初始化AlibabaHealthVaccinNoticeTimebucketRemindAPIRequest对象
func NewAlibabaHealthVaccinNoticeTimebucketRemindRequest() *AlibabaHealthVaccinNoticeTimebucketRemindAPIRequest {
	return &AlibabaHealthVaccinNoticeTimebucketRemindAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaHealthVaccinNoticeTimebucketRemindAPIRequest) GetApiMethodName() string {
	return "alibaba.health.vaccin.notice.timebucket.remind"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaHealthVaccinNoticeTimebucketRemindAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is AlipayUserId Setter
// 432421
func (r *AlibabaHealthVaccinNoticeTimebucketRemindAPIRequest) SetAlipayUserId(_alipayUserId string) error {
	r._alipayUserId = _alipayUserId
	r.Set("alipay_user_id", _alipayUserId)
	return nil
}

// Get AlipayUserId Getter
func (r AlibabaHealthVaccinNoticeTimebucketRemindAPIRequest) GetAlipayUserId() string {
	return r._alipayUserId
}

// Set is VaccineName Setter
// 疫苗名称
func (r *AlibabaHealthVaccinNoticeTimebucketRemindAPIRequest) SetVaccineName(_vaccineName string) error {
	r._vaccineName = _vaccineName
	r.Set("vaccine_name", _vaccineName)
	return nil
}

// Get VaccineName Getter
func (r AlibabaHealthVaccinNoticeTimebucketRemindAPIRequest) GetVaccineName() string {
	return r._vaccineName
}

// Set is ReserveDate Setter
// 预约日期：2019-02-08 严格按照
func (r *AlibabaHealthVaccinNoticeTimebucketRemindAPIRequest) SetReserveDate(_reserveDate string) error {
	r._reserveDate = _reserveDate
	r.Set("reserve_date", _reserveDate)
	return nil
}

// Get ReserveDate Getter
func (r AlibabaHealthVaccinNoticeTimebucketRemindAPIRequest) GetReserveDate() string {
	return r._reserveDate
}

// Set is Name Setter
// 接种人姓名
func (r *AlibabaHealthVaccinNoticeTimebucketRemindAPIRequest) SetName(_name string) error {
	r._name = _name
	r.Set("name", _name)
	return nil
}

// Get Name Getter
func (r AlibabaHealthVaccinNoticeTimebucketRemindAPIRequest) GetName() string {
	return r._name
}

// Set is TheTimes Setter
// 针次
func (r *AlibabaHealthVaccinNoticeTimebucketRemindAPIRequest) SetTheTimes(_theTimes string) error {
	r._theTimes = _theTimes
	r.Set("the_times", _theTimes)
	return nil
}

// Get TheTimes Getter
func (r AlibabaHealthVaccinNoticeTimebucketRemindAPIRequest) GetTheTimes() string {
	return r._theTimes
}

// Set is PovStoreName Setter
// 接种点名称（通知方）
func (r *AlibabaHealthVaccinNoticeTimebucketRemindAPIRequest) SetPovStoreName(_povStoreName string) error {
	r._povStoreName = _povStoreName
	r.Set("pov_store_name", _povStoreName)
	return nil
}

// Get PovStoreName Getter
func (r AlibabaHealthVaccinNoticeTimebucketRemindAPIRequest) GetPovStoreName() string {
	return r._povStoreName
}

// Set is ReserveTime Setter
// 可预约时段
func (r *AlibabaHealthVaccinNoticeTimebucketRemindAPIRequest) SetReserveTime(_reserveTime string) error {
	r._reserveTime = _reserveTime
	r.Set("reserve_time", _reserveTime)
	return nil
}

// Get ReserveTime Getter
func (r AlibabaHealthVaccinNoticeTimebucketRemindAPIRequest) GetReserveTime() string {
	return r._reserveTime
}

// Set is Mobile Setter
// 用户授权的手机号
func (r *AlibabaHealthVaccinNoticeTimebucketRemindAPIRequest) SetMobile(_mobile string) error {
	r._mobile = _mobile
	r.Set("mobile", _mobile)
	return nil
}

// Get Mobile Getter
func (r AlibabaHealthVaccinNoticeTimebucketRemindAPIRequest) GetMobile() string {
	return r._mobile
}
