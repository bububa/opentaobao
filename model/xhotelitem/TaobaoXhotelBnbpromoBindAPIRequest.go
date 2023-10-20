package xhotelitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoxhotelbnbpromobindAPIRequest 自促活动绑定接口 API请求
// taobao.xhotel.bnbpromo.bind
//
// 自促活动绑定接口
type TaobaoxhotelbnbpromobindAPIRequest struct {
	model.Params
	// 营销活动code
	_activityCode string
	// 活动入住时间，民宿通用营销必填
	_checkInFrom string
	// 活动离店时间，民宿通用营销必填
	_checkOutTo string
	// 外部价格信息
	_rateInfos *PromoRateInfo
}

// NewTaobaoxhotelbnbpromobindRequest 初始化TaobaoxhotelbnbpromobindAPIRequest对象
func NewTaobaoxhotelbnbpromobindRequest() *TaobaoxhotelbnbpromobindAPIRequest {
	return &TaobaoxhotelbnbpromobindAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoxhotelbnbpromobindAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.bnbpromo.bind"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoxhotelbnbpromobindAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoxhotelbnbpromobindAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetActivityCode is ActivityCode Setter
// 营销活动code
func (r *TaobaoxhotelbnbpromobindAPIRequest) SetActivityCode(_activityCode string) error {
	r._activityCode = _activityCode
	r.Set("activity_code", _activityCode)
	return nil
}

// GetActivityCode ActivityCode Getter
func (r TaobaoxhotelbnbpromobindAPIRequest) GetActivityCode() string {
	return r._activityCode
}

// SetCheckInFrom is CheckInFrom Setter
// 活动入住时间，民宿通用营销必填
func (r *TaobaoxhotelbnbpromobindAPIRequest) SetCheckInFrom(_checkInFrom string) error {
	r._checkInFrom = _checkInFrom
	r.Set("check_in_from", _checkInFrom)
	return nil
}

// GetCheckInFrom CheckInFrom Getter
func (r TaobaoxhotelbnbpromobindAPIRequest) GetCheckInFrom() string {
	return r._checkInFrom
}

// SetCheckOutTo is CheckOutTo Setter
// 活动离店时间，民宿通用营销必填
func (r *TaobaoxhotelbnbpromobindAPIRequest) SetCheckOutTo(_checkOutTo string) error {
	r._checkOutTo = _checkOutTo
	r.Set("check_out_to", _checkOutTo)
	return nil
}

// GetCheckOutTo CheckOutTo Getter
func (r TaobaoxhotelbnbpromobindAPIRequest) GetCheckOutTo() string {
	return r._checkOutTo
}

// SetRateInfos is RateInfos Setter
// 外部价格信息
func (r *TaobaoxhotelbnbpromobindAPIRequest) SetRateInfos(_rateInfos *PromoRateInfo) error {
	r._rateInfos = _rateInfos
	r.Set("rate_infos", _rateInfos)
	return nil
}

// GetRateInfos RateInfos Getter
func (r TaobaoxhotelbnbpromobindAPIRequest) GetRateInfos() *PromoRateInfo {
	return r._rateInfos
}
