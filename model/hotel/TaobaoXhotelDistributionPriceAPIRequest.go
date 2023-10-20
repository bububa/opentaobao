package hotel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoxhoteldistributionpriceAPIRequest 飞猪分销通用酒店报价接口 API请求
// taobao.xhotel.distribution.price
//
// 飞猪分销通用酒店报价接口
type TaobaoxhoteldistributionpriceAPIRequest struct {
	model.Params
	// 入住日期 yyyy-MM-dd
	_checkinDate string
	// 离店日期 yyyy-MM-dd
	_checkoutDate string
	// 日历报价入住开始日期 yyyy-MM-dd
	_calendarCheckinStartDate string
	// 日历报价入住结束日期 yyyy-MM-dd
	_calendarCheckinEndDate string
	// 查询报价的酒店列表
	_shids int64
	// 是否日历报价计算，false只用填入离日期，true只用填日历开始结束日期
	_isCalendar bool
}

// NewTaobaoxhoteldistributionpriceRequest 初始化TaobaoxhoteldistributionpriceAPIRequest对象
func NewTaobaoxhoteldistributionpriceRequest() *TaobaoxhoteldistributionpriceAPIRequest {
	return &TaobaoxhoteldistributionpriceAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoxhoteldistributionpriceAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.distribution.price"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoxhoteldistributionpriceAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoxhoteldistributionpriceAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCheckinDate is CheckinDate Setter
// 入住日期 yyyy-MM-dd
func (r *TaobaoxhoteldistributionpriceAPIRequest) SetCheckinDate(_checkinDate string) error {
	r._checkinDate = _checkinDate
	r.Set("checkin_date", _checkinDate)
	return nil
}

// GetCheckinDate CheckinDate Getter
func (r TaobaoxhoteldistributionpriceAPIRequest) GetCheckinDate() string {
	return r._checkinDate
}

// SetCheckoutDate is CheckoutDate Setter
// 离店日期 yyyy-MM-dd
func (r *TaobaoxhoteldistributionpriceAPIRequest) SetCheckoutDate(_checkoutDate string) error {
	r._checkoutDate = _checkoutDate
	r.Set("checkout_date", _checkoutDate)
	return nil
}

// GetCheckoutDate CheckoutDate Getter
func (r TaobaoxhoteldistributionpriceAPIRequest) GetCheckoutDate() string {
	return r._checkoutDate
}

// SetCalendarCheckinStartDate is CalendarCheckinStartDate Setter
// 日历报价入住开始日期 yyyy-MM-dd
func (r *TaobaoxhoteldistributionpriceAPIRequest) SetCalendarCheckinStartDate(_calendarCheckinStartDate string) error {
	r._calendarCheckinStartDate = _calendarCheckinStartDate
	r.Set("calendar_checkin_start_date", _calendarCheckinStartDate)
	return nil
}

// GetCalendarCheckinStartDate CalendarCheckinStartDate Getter
func (r TaobaoxhoteldistributionpriceAPIRequest) GetCalendarCheckinStartDate() string {
	return r._calendarCheckinStartDate
}

// SetCalendarCheckinEndDate is CalendarCheckinEndDate Setter
// 日历报价入住结束日期 yyyy-MM-dd
func (r *TaobaoxhoteldistributionpriceAPIRequest) SetCalendarCheckinEndDate(_calendarCheckinEndDate string) error {
	r._calendarCheckinEndDate = _calendarCheckinEndDate
	r.Set("calendar_checkin_end_date", _calendarCheckinEndDate)
	return nil
}

// GetCalendarCheckinEndDate CalendarCheckinEndDate Getter
func (r TaobaoxhoteldistributionpriceAPIRequest) GetCalendarCheckinEndDate() string {
	return r._calendarCheckinEndDate
}

// SetShids is Shids Setter
// 查询报价的酒店列表
func (r *TaobaoxhoteldistributionpriceAPIRequest) SetShids(_shids int64) error {
	r._shids = _shids
	r.Set("shids", _shids)
	return nil
}

// GetShids Shids Getter
func (r TaobaoxhoteldistributionpriceAPIRequest) GetShids() int64 {
	return r._shids
}

// SetIsCalendar is IsCalendar Setter
// 是否日历报价计算，false只用填入离日期，true只用填日历开始结束日期
func (r *TaobaoxhoteldistributionpriceAPIRequest) SetIsCalendar(_isCalendar bool) error {
	r._isCalendar = _isCalendar
	r.Set("is_calendar", _isCalendar)
	return nil
}

// GetIsCalendar IsCalendar Getter
func (r TaobaoxhoteldistributionpriceAPIRequest) GetIsCalendar() bool {
	return r._isCalendar
}
