package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallmallitemcentersupplierabilityupdateAPIRequest 门店服务能力授权接口 API请求
// tmall.mallitemcenter.supplier.ability.update
//
// 门店服务能力授权
type TmallmallitemcentersupplierabilityupdateAPIRequest struct {
	model.Params
	// 入参
	_param0 *EnableServiceStoreRequestDto
}

// NewTmallmallitemcentersupplierabilityupdateRequest 初始化TmallmallitemcentersupplierabilityupdateAPIRequest对象
func NewTmallmallitemcentersupplierabilityupdateRequest() *TmallmallitemcentersupplierabilityupdateAPIRequest {
	return &TmallmallitemcentersupplierabilityupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallmallitemcentersupplierabilityupdateAPIRequest) GetApiMethodName() string {
	return "tmall.mallitemcenter.supplier.ability.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallmallitemcentersupplierabilityupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallmallitemcentersupplierabilityupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 入参
func (r *TmallmallitemcentersupplierabilityupdateAPIRequest) SetParam0(_param0 *EnableServiceStoreRequestDto) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r TmallmallitemcentersupplierabilityupdateAPIRequest) GetParam0() *EnableServiceStoreRequestDto {
	return r._param0
}
