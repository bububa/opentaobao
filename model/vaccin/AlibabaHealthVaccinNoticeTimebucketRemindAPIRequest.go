package vaccin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabahealthvaccinnoticetimebucketremindAPIRequest 疫苗预约时间段提醒 API请求
// alibaba.health.vaccin.notice.timebucket.remind
//
// 疫苗预约时间段提醒
type AlibabahealthvaccinnoticetimebucketremindAPIRequest struct {
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

// NewAlibabahealthvaccinnoticetimebucketremindRequest 初始化AlibabahealthvaccinnoticetimebucketremindAPIRequest对象
func NewAlibabahealthvaccinnoticetimebucketremindRequest() *AlibabahealthvaccinnoticetimebucketremindAPIRequest {
	return &AlibabahealthvaccinnoticetimebucketremindAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabahealthvaccinnoticetimebucketremindAPIRequest) GetApiMethodName() string {
	return "alibaba.health.vaccin.notice.timebucket.remind"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabahealthvaccinnoticetimebucketremindAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabahealthvaccinnoticetimebucketremindAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAlipayUserId is AlipayUserId Setter
// 432421
func (r *AlibabahealthvaccinnoticetimebucketremindAPIRequest) SetAlipayUserId(_alipayUserId string) error {
	r._alipayUserId = _alipayUserId
	r.Set("alipay_user_id", _alipayUserId)
	return nil
}

// GetAlipayUserId AlipayUserId Getter
func (r AlibabahealthvaccinnoticetimebucketremindAPIRequest) GetAlipayUserId() string {
	return r._alipayUserId
}

// SetVaccineName is VaccineName Setter
// 疫苗名称
func (r *AlibabahealthvaccinnoticetimebucketremindAPIRequest) SetVaccineName(_vaccineName string) error {
	r._vaccineName = _vaccineName
	r.Set("vaccine_name", _vaccineName)
	return nil
}

// GetVaccineName VaccineName Getter
func (r AlibabahealthvaccinnoticetimebucketremindAPIRequest) GetVaccineName() string {
	return r._vaccineName
}

// SetReserveDate is ReserveDate Setter
// 预约日期：2019-02-08 严格按照
func (r *AlibabahealthvaccinnoticetimebucketremindAPIRequest) SetReserveDate(_reserveDate string) error {
	r._reserveDate = _reserveDate
	r.Set("reserve_date", _reserveDate)
	return nil
}

// GetReserveDate ReserveDate Getter
func (r AlibabahealthvaccinnoticetimebucketremindAPIRequest) GetReserveDate() string {
	return r._reserveDate
}

// SetName is Name Setter
// 接种人姓名
func (r *AlibabahealthvaccinnoticetimebucketremindAPIRequest) SetName(_name string) error {
	r._name = _name
	r.Set("name", _name)
	return nil
}

// GetName Name Getter
func (r AlibabahealthvaccinnoticetimebucketremindAPIRequest) GetName() string {
	return r._name
}

// SetTheTimes is TheTimes Setter
// 针次
func (r *AlibabahealthvaccinnoticetimebucketremindAPIRequest) SetTheTimes(_theTimes string) error {
	r._theTimes = _theTimes
	r.Set("the_times", _theTimes)
	return nil
}

// GetTheTimes TheTimes Getter
func (r AlibabahealthvaccinnoticetimebucketremindAPIRequest) GetTheTimes() string {
	return r._theTimes
}

// SetPovStoreName is PovStoreName Setter
// 接种点名称（通知方）
func (r *AlibabahealthvaccinnoticetimebucketremindAPIRequest) SetPovStoreName(_povStoreName string) error {
	r._povStoreName = _povStoreName
	r.Set("pov_store_name", _povStoreName)
	return nil
}

// GetPovStoreName PovStoreName Getter
func (r AlibabahealthvaccinnoticetimebucketremindAPIRequest) GetPovStoreName() string {
	return r._povStoreName
}

// SetReserveTime is ReserveTime Setter
// 可预约时段
func (r *AlibabahealthvaccinnoticetimebucketremindAPIRequest) SetReserveTime(_reserveTime string) error {
	r._reserveTime = _reserveTime
	r.Set("reserve_time", _reserveTime)
	return nil
}

// GetReserveTime ReserveTime Getter
func (r AlibabahealthvaccinnoticetimebucketremindAPIRequest) GetReserveTime() string {
	return r._reserveTime
}

// SetMobile is Mobile Setter
// 用户授权的手机号
func (r *AlibabahealthvaccinnoticetimebucketremindAPIRequest) SetMobile(_mobile string) error {
	r._mobile = _mobile
	r.Set("mobile", _mobile)
	return nil
}

// GetMobile Mobile Getter
func (r AlibabahealthvaccinnoticetimebucketremindAPIRequest) GetMobile() string {
	return r._mobile
}
