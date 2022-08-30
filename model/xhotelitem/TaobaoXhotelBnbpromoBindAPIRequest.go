package xhotelitem

import (
	"net/url"

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
	// 外部价格信息
	_rateInfos *PromoRateInfo
}

// NewTaobaoXhotelBnbpromoBindRequest 初始化TaobaoXhotelBnbpromoBindAPIRequest对象
func NewTaobaoXhotelBnbpromoBindRequest() *TaobaoXhotelBnbpromoBindAPIRequest {
	return &TaobaoXhotelBnbpromoBindAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoXhotelBnbpromoBindAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.bnbpromo.bind"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoXhotelBnbpromoBindAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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
