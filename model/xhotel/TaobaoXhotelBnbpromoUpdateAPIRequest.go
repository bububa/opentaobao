package xhotel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelBnbpromoUpdateAPIRequest 民宿营销活动更新 API请求
// taobao.xhotel.bnbpromo.update
//
// 全量更新对应外部活动code相关的营销活动信息
type TaobaoXhotelBnbpromoUpdateAPIRequest struct {
	model.Params
	// 更新营销活动的入参
	_updatePromoParam *UpdatePromoParam
}

// NewTaobaoXhotelBnbpromoUpdateRequest 初始化TaobaoXhotelBnbpromoUpdateAPIRequest对象
func NewTaobaoXhotelBnbpromoUpdateRequest() *TaobaoXhotelBnbpromoUpdateAPIRequest {
	return &TaobaoXhotelBnbpromoUpdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoXhotelBnbpromoUpdateAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.bnbpromo.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoXhotelBnbpromoUpdateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetUpdatePromoParam is UpdatePromoParam Setter
// 更新营销活动的入参
func (r *TaobaoXhotelBnbpromoUpdateAPIRequest) SetUpdatePromoParam(_updatePromoParam *UpdatePromoParam) error {
	r._updatePromoParam = _updatePromoParam
	r.Set("update_promo_param", _updatePromoParam)
	return nil
}

// GetUpdatePromoParam UpdatePromoParam Getter
func (r TaobaoXhotelBnbpromoUpdateAPIRequest) GetUpdatePromoParam() *UpdatePromoParam {
	return r._updatePromoParam
}
