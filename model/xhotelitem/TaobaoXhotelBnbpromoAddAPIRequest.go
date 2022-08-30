package xhotelitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelBnbpromoAddAPIRequest 自促活动申请接口 API请求
// taobao.xhotel.bnbpromo.add
//
// 自促活动申请接口
type TaobaoXhotelBnbpromoAddAPIRequest struct {
	model.Params
	// 营销类型
	_promoInfo *PromoInfo
}

// NewTaobaoXhotelBnbpromoAddRequest 初始化TaobaoXhotelBnbpromoAddAPIRequest对象
func NewTaobaoXhotelBnbpromoAddRequest() *TaobaoXhotelBnbpromoAddAPIRequest {
	return &TaobaoXhotelBnbpromoAddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoXhotelBnbpromoAddAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.bnbpromo.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoXhotelBnbpromoAddAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetPromoInfo is PromoInfo Setter
// 营销类型
func (r *TaobaoXhotelBnbpromoAddAPIRequest) SetPromoInfo(_promoInfo *PromoInfo) error {
	r._promoInfo = _promoInfo
	r.Set("promo_info", _promoInfo)
	return nil
}

// GetPromoInfo PromoInfo Getter
func (r TaobaoXhotelBnbpromoAddAPIRequest) GetPromoInfo() *PromoInfo {
	return r._promoInfo
}
