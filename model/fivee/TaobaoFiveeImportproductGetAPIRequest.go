package fivee

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFiveeImportproductGetAPIRequest 进口商品查询 API请求
// taobao.fivee.importproduct.get
//
// 资质共享平台查询进口商品信息
type TaobaoFiveeImportproductGetAPIRequest struct {
	model.Params
	// bu身份标识
	_paramBuCode string
	// 条形码
	_paramBarcode string
}

// NewTaobaoFiveeImportproductGetRequest 初始化TaobaoFiveeImportproductGetAPIRequest对象
func NewTaobaoFiveeImportproductGetRequest() *TaobaoFiveeImportproductGetAPIRequest {
	return &TaobaoFiveeImportproductGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoFiveeImportproductGetAPIRequest) GetApiMethodName() string {
	return "taobao.fivee.importproduct.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoFiveeImportproductGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ParamBuCode Setter
// bu身份标识
func (r *TaobaoFiveeImportproductGetAPIRequest) SetParamBuCode(_paramBuCode string) error {
	r._paramBuCode = _paramBuCode
	r.Set("param_bu_code", _paramBuCode)
	return nil
}

// Get ParamBuCode Getter
func (r TaobaoFiveeImportproductGetAPIRequest) GetParamBuCode() string {
	return r._paramBuCode
}

// Set is ParamBarcode Setter
// 条形码
func (r *TaobaoFiveeImportproductGetAPIRequest) SetParamBarcode(_paramBarcode string) error {
	r._paramBarcode = _paramBarcode
	r.Set("param_barcode", _paramBarcode)
	return nil
}

// Get ParamBarcode Getter
func (r TaobaoFiveeImportproductGetAPIRequest) GetParamBarcode() string {
	return r._paramBarcode
}
