package fivee

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFiveeInnerproductGetAPIRequest 国产商品资质查询 API请求
// taobao.fivee.innerproduct.get
//
// 资质共享平台，国产商品查询
type TaobaoFiveeInnerproductGetAPIRequest struct {
	model.Params
	// bu身份标识
	_paramBucode string
	// 条形码
	_paramBarcode string
}

// NewTaobaoFiveeInnerproductGetRequest 初始化TaobaoFiveeInnerproductGetAPIRequest对象
func NewTaobaoFiveeInnerproductGetRequest() *TaobaoFiveeInnerproductGetAPIRequest {
	return &TaobaoFiveeInnerproductGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoFiveeInnerproductGetAPIRequest) GetApiMethodName() string {
	return "taobao.fivee.innerproduct.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoFiveeInnerproductGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoFiveeInnerproductGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamBucode is ParamBucode Setter
// bu身份标识
func (r *TaobaoFiveeInnerproductGetAPIRequest) SetParamBucode(_paramBucode string) error {
	r._paramBucode = _paramBucode
	r.Set("param_bucode", _paramBucode)
	return nil
}

// GetParamBucode ParamBucode Getter
func (r TaobaoFiveeInnerproductGetAPIRequest) GetParamBucode() string {
	return r._paramBucode
}

// SetParamBarcode is ParamBarcode Setter
// 条形码
func (r *TaobaoFiveeInnerproductGetAPIRequest) SetParamBarcode(_paramBarcode string) error {
	r._paramBarcode = _paramBarcode
	r.Set("param_barcode", _paramBarcode)
	return nil
}

// GetParamBarcode ParamBarcode Getter
func (r TaobaoFiveeInnerproductGetAPIRequest) GetParamBarcode() string {
	return r._paramBarcode
}
