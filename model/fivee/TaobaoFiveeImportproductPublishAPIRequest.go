package fivee

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFiveeImportproductPublishAPIRequest 进口商品发布 API请求
// taobao.fivee.importproduct.publish
//
// 直营业务商家入住发布商品时，上传商品及商家证照信息
type TaobaoFiveeImportproductPublishAPIRequest struct {
	model.Params
	// 进口商品
	_importProduct *ImportProduct
	// bu身份标识
	_paramBucode string
}

// NewTaobaoFiveeImportproductPublishRequest 初始化TaobaoFiveeImportproductPublishAPIRequest对象
func NewTaobaoFiveeImportproductPublishRequest() *TaobaoFiveeImportproductPublishAPIRequest {
	return &TaobaoFiveeImportproductPublishAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoFiveeImportproductPublishAPIRequest) GetApiMethodName() string {
	return "taobao.fivee.importproduct.publish"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoFiveeImportproductPublishAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ImportProduct Setter
// 进口商品
func (r *TaobaoFiveeImportproductPublishAPIRequest) SetImportProduct(_importProduct *ImportProduct) error {
	r._importProduct = _importProduct
	r.Set("import_product", _importProduct)
	return nil
}

// Get ImportProduct Getter
func (r TaobaoFiveeImportproductPublishAPIRequest) GetImportProduct() *ImportProduct {
	return r._importProduct
}

// Set is ParamBucode Setter
// bu身份标识
func (r *TaobaoFiveeImportproductPublishAPIRequest) SetParamBucode(_paramBucode string) error {
	r._paramBucode = _paramBucode
	r.Set("param_bucode", _paramBucode)
	return nil
}

// Get ParamBucode Getter
func (r TaobaoFiveeImportproductPublishAPIRequest) GetParamBucode() string {
	return r._paramBucode
}
