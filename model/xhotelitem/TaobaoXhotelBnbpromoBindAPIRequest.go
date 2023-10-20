package xhotelitem

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelBnbpromoBindAPIRequest 自促活动绑定接口 API请求
// taobao.xhotel.bnbpromo.bind
//
// 自促活动绑定接口
type TaobaoXhotelBnbpromoBindAPIRequest struct {
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

// NewTaobaoXhotelBnbpromoBindRequest 初始化TaobaoXhotelBnbpromoBindAPIRequest对象
func NewTaobaoXhotelBnbpromoBindRequest() *TaobaoXhotelBnbpromoBindAPIRequest {
	return &TaobaoXhotelBnbpromoBindAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoXhotelBnbpromoBindAPIRequest) Reset() {
	r._activityCode = ""
	r._checkInFrom = ""
	r._checkOutTo = ""
	r._rateInfos = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoXhotelBnbpromoBindAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.bnbpromo.bind"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoXhotelBnbpromoBindAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoXhotelBnbpromoBindAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetActivityCode is ActivityCode Setter
// 营销活动code
func (r *TaobaoXhotelBnbpromoBindAPIRequest) SetActivityCode(_activityCode string) error {
	r._activityCode = _activityCode
	r.Set("activity_code", _activityCode)
	return nil
}

// GetActivityCode ActivityCode Getter
func (r TaobaoXhotelBnbpromoBindAPIRequest) GetActivityCode() string {
	return r._activityCode
}

// SetCheckInFrom is CheckInFrom Setter
// 活动入住时间，民宿通用营销必填
func (r *TaobaoXhotelBnbpromoBindAPIRequest) SetCheckInFrom(_checkInFrom string) error {
	r._checkInFrom = _checkInFrom
	r.Set("check_in_from", _checkInFrom)
	return nil
}

// GetCheckInFrom CheckInFrom Getter
func (r TaobaoXhotelBnbpromoBindAPIRequest) GetCheckInFrom() string {
	return r._checkInFrom
}

// SetCheckOutTo is CheckOutTo Setter
// 活动离店时间，民宿通用营销必填
func (r *TaobaoXhotelBnbpromoBindAPIRequest) SetCheckOutTo(_checkOutTo string) error {
	r._checkOutTo = _checkOutTo
	r.Set("check_out_to", _checkOutTo)
	return nil
}

// GetCheckOutTo CheckOutTo Getter
func (r TaobaoXhotelBnbpromoBindAPIRequest) GetCheckOutTo() string {
	return r._checkOutTo
}

// SetRateInfos is RateInfos Setter
// 外部价格信息
func (r *TaobaoXhotelBnbpromoBindAPIRequest) SetRateInfos(_rateInfos *PromoRateInfo) error {
	r._rateInfos = _rateInfos
	r.Set("rate_infos", _rateInfos)
	return nil
}

// GetRateInfos RateInfos Getter
func (r TaobaoXhotelBnbpromoBindAPIRequest) GetRateInfos() *PromoRateInfo {
	return r._rateInfos
}

var poolTaobaoXhotelBnbpromoBindAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoXhotelBnbpromoBindRequest()
	},
}

// GetTaobaoXhotelBnbpromoBindRequest 从 sync.Pool 获取 TaobaoXhotelBnbpromoBindAPIRequest
func GetTaobaoXhotelBnbpromoBindAPIRequest() *TaobaoXhotelBnbpromoBindAPIRequest {
	return poolTaobaoXhotelBnbpromoBindAPIRequest.Get().(*TaobaoXhotelBnbpromoBindAPIRequest)
}

// ReleaseTaobaoXhotelBnbpromoBindAPIRequest 将 TaobaoXhotelBnbpromoBindAPIRequest 放入 sync.Pool
func ReleaseTaobaoXhotelBnbpromoBindAPIRequest(v *TaobaoXhotelBnbpromoBindAPIRequest) {
	v.Reset()
	poolTaobaoXhotelBnbpromoBindAPIRequest.Put(v)
}
