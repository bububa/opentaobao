package fivee

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaofiveeinnerproductpublishAPIRequest 国产商品发布 API请求
// taobao.fivee.innerproduct.publish
//
// 资质共享平台国产商品发布
type TaobaofiveeinnerproductpublishAPIRequest struct {
	model.Params
	// bu身份标识
	_paramBucode string
	// 国产商品
	_paramInnerProduct *InnerProduct
}

// NewTaobaofiveeinnerproductpublishRequest 初始化TaobaofiveeinnerproductpublishAPIRequest对象
func NewTaobaofiveeinnerproductpublishRequest() *TaobaofiveeinnerproductpublishAPIRequest {
	return &TaobaofiveeinnerproductpublishAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaofiveeinnerproductpublishAPIRequest) GetApiMethodName() string {
	return "taobao.fivee.innerproduct.publish"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaofiveeinnerproductpublishAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaofiveeinnerproductpublishAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamBucode is ParamBucode Setter
// bu身份标识
func (r *TaobaofiveeinnerproductpublishAPIRequest) SetParamBucode(_paramBucode string) error {
	r._paramBucode = _paramBucode
	r.Set("param_bucode", _paramBucode)
	return nil
}

// GetParamBucode ParamBucode Getter
func (r TaobaofiveeinnerproductpublishAPIRequest) GetParamBucode() string {
	return r._paramBucode
}

// SetParamInnerProduct is ParamInnerProduct Setter
// 国产商品
func (r *TaobaofiveeinnerproductpublishAPIRequest) SetParamInnerProduct(_paramInnerProduct *InnerProduct) error {
	r._paramInnerProduct = _paramInnerProduct
	r.Set("param_inner_product", _paramInnerProduct)
	return nil
}

// GetParamInnerProduct ParamInnerProduct Getter
func (r TaobaofiveeinnerproductpublishAPIRequest) GetParamInnerProduct() *InnerProduct {
	return r._paramInnerProduct
}
