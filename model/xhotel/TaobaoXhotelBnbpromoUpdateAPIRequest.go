package xhotel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoxhotelbnbpromoupdateAPIRequest 民宿营销活动更新 API请求
// taobao.xhotel.bnbpromo.update
//
// 全量更新对应外部活动code相关的营销活动信息
type TaobaoxhotelbnbpromoupdateAPIRequest struct {
	model.Params
	// 更新营销活动的入参
	_updatePromoParam *UpdatePromoParam
}

// NewTaobaoxhotelbnbpromoupdateRequest 初始化TaobaoxhotelbnbpromoupdateAPIRequest对象
func NewTaobaoxhotelbnbpromoupdateRequest() *TaobaoxhotelbnbpromoupdateAPIRequest {
	return &TaobaoxhotelbnbpromoupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoxhotelbnbpromoupdateAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.bnbpromo.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoxhotelbnbpromoupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoxhotelbnbpromoupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUpdatePromoParam is UpdatePromoParam Setter
// 更新营销活动的入参
func (r *TaobaoxhotelbnbpromoupdateAPIRequest) SetUpdatePromoParam(_updatePromoParam *UpdatePromoParam) error {
	r._updatePromoParam = _updatePromoParam
	r.Set("update_promo_param", _updatePromoParam)
	return nil
}

// GetUpdatePromoParam UpdatePromoParam Getter
func (r TaobaoxhotelbnbpromoupdateAPIRequest) GetUpdatePromoParam() *UpdatePromoParam {
	return r._updatePromoParam
}
