package mos

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabamjocwritesaleslipAPIRequest 开票占库 API请求
// alibaba.mj.oc.writesaleslip
//
// 开票占库
type AlibabamjocwritesaleslipAPIRequest struct {
	model.Params
	// 开票占库入参
	_posSaleOrder *PosSaleOrderDto
}

// NewAlibabamjocwritesaleslipRequest 初始化AlibabamjocwritesaleslipAPIRequest对象
func NewAlibabamjocwritesaleslipRequest() *AlibabamjocwritesaleslipAPIRequest {
	return &AlibabamjocwritesaleslipAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabamjocwritesaleslipAPIRequest) GetApiMethodName() string {
	return "alibaba.mj.oc.writesaleslip"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabamjocwritesaleslipAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabamjocwritesaleslipAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPosSaleOrder is PosSaleOrder Setter
// 开票占库入参
func (r *AlibabamjocwritesaleslipAPIRequest) SetPosSaleOrder(_posSaleOrder *PosSaleOrderDto) error {
	r._posSaleOrder = _posSaleOrder
	r.Set("pos_sale_order", _posSaleOrder)
	return nil
}

// GetPosSaleOrder PosSaleOrder Getter
func (r AlibabamjocwritesaleslipAPIRequest) GetPosSaleOrder() *PosSaleOrderDto {
	return r._posSaleOrder
}
