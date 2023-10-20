package xhotelonlineorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoxhotelintlrateupdateAPIRequest 不落库商家推送更新酒店rate API请求
// taobao.xhotel.intl.rate.update
//
// 商家主动推送不落库商品的酒店信息
type TaobaoxhotelintlrateupdateAPIRequest struct {
	model.Params
	// rate更新参数
	_updateRateParam *UpdateRateParam
}

// NewTaobaoxhotelintlrateupdateRequest 初始化TaobaoxhotelintlrateupdateAPIRequest对象
func NewTaobaoxhotelintlrateupdateRequest() *TaobaoxhotelintlrateupdateAPIRequest {
	return &TaobaoxhotelintlrateupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoxhotelintlrateupdateAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.intl.rate.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoxhotelintlrateupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoxhotelintlrateupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUpdateRateParam is UpdateRateParam Setter
// rate更新参数
func (r *TaobaoxhotelintlrateupdateAPIRequest) SetUpdateRateParam(_updateRateParam *UpdateRateParam) error {
	r._updateRateParam = _updateRateParam
	r.Set("update_rate_param", _updateRateParam)
	return nil
}

// GetUpdateRateParam UpdateRateParam Getter
func (r TaobaoxhotelintlrateupdateAPIRequest) GetUpdateRateParam() *UpdateRateParam {
	return r._updateRateParam
}
