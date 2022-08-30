package xhotelitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelBnbpromoUnbindAPIRequest 自促活动解绑接口 API请求
// taobao.xhotel.bnbpromo.unbind
//
// 自促活动解绑接口
type TaobaoXhotelBnbpromoUnbindAPIRequest struct {
	model.Params
	// 营销活动code
	_activityCode string
	// 营销
	_rateInfos *PromoRateInfo
}

// NewTaobaoXhotelBnbpromoUnbindRequest 初始化TaobaoXhotelBnbpromoUnbindAPIRequest对象
func NewTaobaoXhotelBnbpromoUnbindRequest() *TaobaoXhotelBnbpromoUnbindAPIRequest {
	return &TaobaoXhotelBnbpromoUnbindAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoXhotelBnbpromoUnbindAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.bnbpromo.unbind"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoXhotelBnbpromoUnbindAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetActivityCode is ActivityCode Setter
// 营销活动code
func (r *TaobaoXhotelBnbpromoUnbindAPIRequest) SetActivityCode(_activityCode string) error {
	r._activityCode = _activityCode
	r.Set("activity_code", _activityCode)
	return nil
}

// GetActivityCode ActivityCode Getter
func (r TaobaoXhotelBnbpromoUnbindAPIRequest) GetActivityCode() string {
	return r._activityCode
}

// SetRateInfos is RateInfos Setter
// 营销
func (r *TaobaoXhotelBnbpromoUnbindAPIRequest) SetRateInfos(_rateInfos *PromoRateInfo) error {
	r._rateInfos = _rateInfos
	r.Set("rate_infos", _rateInfos)
	return nil
}

// GetRateInfos RateInfos Getter
func (r TaobaoXhotelBnbpromoUnbindAPIRequest) GetRateInfos() *PromoRateInfo {
	return r._rateInfos
}
