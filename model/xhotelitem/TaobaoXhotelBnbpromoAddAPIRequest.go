package xhotelitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoxhotelbnbpromoaddAPIRequest 自促活动申请接口 API请求
// taobao.xhotel.bnbpromo.add
//
// 自促活动申请接口
type TaobaoxhotelbnbpromoaddAPIRequest struct {
	model.Params
	// 营销类型
	_promoInfo *PromoInfo
}

// NewTaobaoxhotelbnbpromoaddRequest 初始化TaobaoxhotelbnbpromoaddAPIRequest对象
func NewTaobaoxhotelbnbpromoaddRequest() *TaobaoxhotelbnbpromoaddAPIRequest {
	return &TaobaoxhotelbnbpromoaddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoxhotelbnbpromoaddAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.bnbpromo.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoxhotelbnbpromoaddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoxhotelbnbpromoaddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPromoInfo is PromoInfo Setter
// 营销类型
func (r *TaobaoxhotelbnbpromoaddAPIRequest) SetPromoInfo(_promoInfo *PromoInfo) error {
	r._promoInfo = _promoInfo
	r.Set("promo_info", _promoInfo)
	return nil
}

// GetPromoInfo PromoInfo Getter
func (r TaobaoxhotelbnbpromoaddAPIRequest) GetPromoInfo() *PromoInfo {
	return r._promoInfo
}
