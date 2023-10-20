package xhotelitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoxhotelbnbpromounbindAPIRequest 自促活动解绑接口 API请求
// taobao.xhotel.bnbpromo.unbind
//
// 自促活动解绑接口
type TaobaoxhotelbnbpromounbindAPIRequest struct {
	model.Params
	// 营销活动code
	_activityCode string
	// 营销
	_rateInfos *PromoRateInfo
}

// NewTaobaoxhotelbnbpromounbindRequest 初始化TaobaoxhotelbnbpromounbindAPIRequest对象
func NewTaobaoxhotelbnbpromounbindRequest() *TaobaoxhotelbnbpromounbindAPIRequest {
	return &TaobaoxhotelbnbpromounbindAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoxhotelbnbpromounbindAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.bnbpromo.unbind"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoxhotelbnbpromounbindAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoxhotelbnbpromounbindAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetActivityCode is ActivityCode Setter
// 营销活动code
func (r *TaobaoxhotelbnbpromounbindAPIRequest) SetActivityCode(_activityCode string) error {
	r._activityCode = _activityCode
	r.Set("activity_code", _activityCode)
	return nil
}

// GetActivityCode ActivityCode Getter
func (r TaobaoxhotelbnbpromounbindAPIRequest) GetActivityCode() string {
	return r._activityCode
}

// SetRateInfos is RateInfos Setter
// 营销
func (r *TaobaoxhotelbnbpromounbindAPIRequest) SetRateInfos(_rateInfos *PromoRateInfo) error {
	r._rateInfos = _rateInfos
	r.Set("rate_infos", _rateInfos)
	return nil
}

// GetRateInfos RateInfos Getter
func (r TaobaoxhotelbnbpromounbindAPIRequest) GetRateInfos() *PromoRateInfo {
	return r._rateInfos
}
