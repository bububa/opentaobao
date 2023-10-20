package fivee

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFiveeInnerproductPublishAPIRequest 国产商品发布 API请求
// taobao.fivee.innerproduct.publish
//
// 资质共享平台国产商品发布
type TaobaoFiveeInnerproductPublishAPIRequest struct {
	model.Params
	// bu身份标识
	_paramBucode string
	// 国产商品
	_paramInnerProduct *InnerProduct
}

// NewTaobaoFiveeInnerproductPublishRequest 初始化TaobaoFiveeInnerproductPublishAPIRequest对象
func NewTaobaoFiveeInnerproductPublishRequest() *TaobaoFiveeInnerproductPublishAPIRequest {
	return &TaobaoFiveeInnerproductPublishAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoFiveeInnerproductPublishAPIRequest) Reset() {
	r._paramBucode = ""
	r._paramInnerProduct = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoFiveeInnerproductPublishAPIRequest) GetApiMethodName() string {
	return "taobao.fivee.innerproduct.publish"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoFiveeInnerproductPublishAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoFiveeInnerproductPublishAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamBucode is ParamBucode Setter
// bu身份标识
func (r *TaobaoFiveeInnerproductPublishAPIRequest) SetParamBucode(_paramBucode string) error {
	r._paramBucode = _paramBucode
	r.Set("param_bucode", _paramBucode)
	return nil
}

// GetParamBucode ParamBucode Getter
func (r TaobaoFiveeInnerproductPublishAPIRequest) GetParamBucode() string {
	return r._paramBucode
}

// SetParamInnerProduct is ParamInnerProduct Setter
// 国产商品
func (r *TaobaoFiveeInnerproductPublishAPIRequest) SetParamInnerProduct(_paramInnerProduct *InnerProduct) error {
	r._paramInnerProduct = _paramInnerProduct
	r.Set("param_inner_product", _paramInnerProduct)
	return nil
}

// GetParamInnerProduct ParamInnerProduct Getter
func (r TaobaoFiveeInnerproductPublishAPIRequest) GetParamInnerProduct() *InnerProduct {
	return r._paramInnerProduct
}

var poolTaobaoFiveeInnerproductPublishAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoFiveeInnerproductPublishRequest()
	},
}

// GetTaobaoFiveeInnerproductPublishRequest 从 sync.Pool 获取 TaobaoFiveeInnerproductPublishAPIRequest
func GetTaobaoFiveeInnerproductPublishAPIRequest() *TaobaoFiveeInnerproductPublishAPIRequest {
	return poolTaobaoFiveeInnerproductPublishAPIRequest.Get().(*TaobaoFiveeInnerproductPublishAPIRequest)
}

// ReleaseTaobaoFiveeInnerproductPublishAPIRequest 将 TaobaoFiveeInnerproductPublishAPIRequest 放入 sync.Pool
func ReleaseTaobaoFiveeInnerproductPublishAPIRequest(v *TaobaoFiveeInnerproductPublishAPIRequest) {
	v.Reset()
	poolTaobaoFiveeInnerproductPublishAPIRequest.Put(v)
}
