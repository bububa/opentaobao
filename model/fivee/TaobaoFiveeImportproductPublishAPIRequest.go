package fivee

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaofiveeimportproductpublishAPIRequest 进口商品发布 API请求
// taobao.fivee.importproduct.publish
//
// 直营业务商家入住发布商品时，上传商品及商家证照信息
type TaobaofiveeimportproductpublishAPIRequest struct {
	model.Params
	// bu身份标识
	_paramBucode string
	// 进口商品
	_importProduct *ImportProduct
}

// NewTaobaofiveeimportproductpublishRequest 初始化TaobaofiveeimportproductpublishAPIRequest对象
func NewTaobaofiveeimportproductpublishRequest() *TaobaofiveeimportproductpublishAPIRequest {
	return &TaobaofiveeimportproductpublishAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaofiveeimportproductpublishAPIRequest) GetApiMethodName() string {
	return "taobao.fivee.importproduct.publish"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaofiveeimportproductpublishAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaofiveeimportproductpublishAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamBucode is ParamBucode Setter
// bu身份标识
func (r *TaobaofiveeimportproductpublishAPIRequest) SetParamBucode(_paramBucode string) error {
	r._paramBucode = _paramBucode
	r.Set("param_bucode", _paramBucode)
	return nil
}

// GetParamBucode ParamBucode Getter
func (r TaobaofiveeimportproductpublishAPIRequest) GetParamBucode() string {
	return r._paramBucode
}

// SetImportProduct is ImportProduct Setter
// 进口商品
func (r *TaobaofiveeimportproductpublishAPIRequest) SetImportProduct(_importProduct *ImportProduct) error {
	r._importProduct = _importProduct
	r.Set("import_product", _importProduct)
	return nil
}

// GetImportProduct ImportProduct Getter
func (r TaobaofiveeimportproductpublishAPIRequest) GetImportProduct() *ImportProduct {
	return r._importProduct
}
