package fivee

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaofiveeimportproductgetAPIRequest 进口商品查询 API请求
// taobao.fivee.importproduct.get
//
// 资质共享平台查询进口商品信息
type TaobaofiveeimportproductgetAPIRequest struct {
	model.Params
	// bu身份标识
	_paramBuCode string
	// 条形码
	_paramBarcode string
}

// NewTaobaofiveeimportproductgetRequest 初始化TaobaofiveeimportproductgetAPIRequest对象
func NewTaobaofiveeimportproductgetRequest() *TaobaofiveeimportproductgetAPIRequest {
	return &TaobaofiveeimportproductgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaofiveeimportproductgetAPIRequest) GetApiMethodName() string {
	return "taobao.fivee.importproduct.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaofiveeimportproductgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaofiveeimportproductgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamBuCode is ParamBuCode Setter
// bu身份标识
func (r *TaobaofiveeimportproductgetAPIRequest) SetParamBuCode(_paramBuCode string) error {
	r._paramBuCode = _paramBuCode
	r.Set("param_bu_code", _paramBuCode)
	return nil
}

// GetParamBuCode ParamBuCode Getter
func (r TaobaofiveeimportproductgetAPIRequest) GetParamBuCode() string {
	return r._paramBuCode
}

// SetParamBarcode is ParamBarcode Setter
// 条形码
func (r *TaobaofiveeimportproductgetAPIRequest) SetParamBarcode(_paramBarcode string) error {
	r._paramBarcode = _paramBarcode
	r.Set("param_barcode", _paramBarcode)
	return nil
}

// GetParamBarcode ParamBarcode Getter
func (r TaobaofiveeimportproductgetAPIRequest) GetParamBarcode() string {
	return r._paramBarcode
}
