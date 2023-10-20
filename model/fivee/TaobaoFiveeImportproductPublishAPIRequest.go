package fivee

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFiveeImportproductPublishAPIRequest 进口商品发布 API请求
// taobao.fivee.importproduct.publish
//
// 直营业务商家入住发布商品时，上传商品及商家证照信息
type TaobaoFiveeImportproductPublishAPIRequest struct {
	model.Params
	// bu身份标识
	_paramBucode string
	// 进口商品
	_importProduct *ImportProduct
}

// NewTaobaoFiveeImportproductPublishRequest 初始化TaobaoFiveeImportproductPublishAPIRequest对象
func NewTaobaoFiveeImportproductPublishRequest() *TaobaoFiveeImportproductPublishAPIRequest {
	return &TaobaoFiveeImportproductPublishAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoFiveeImportproductPublishAPIRequest) Reset() {
	r._paramBucode = ""
	r._importProduct = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoFiveeImportproductPublishAPIRequest) GetApiMethodName() string {
	return "taobao.fivee.importproduct.publish"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoFiveeImportproductPublishAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoFiveeImportproductPublishAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamBucode is ParamBucode Setter
// bu身份标识
func (r *TaobaoFiveeImportproductPublishAPIRequest) SetParamBucode(_paramBucode string) error {
	r._paramBucode = _paramBucode
	r.Set("param_bucode", _paramBucode)
	return nil
}

// GetParamBucode ParamBucode Getter
func (r TaobaoFiveeImportproductPublishAPIRequest) GetParamBucode() string {
	return r._paramBucode
}

// SetImportProduct is ImportProduct Setter
// 进口商品
func (r *TaobaoFiveeImportproductPublishAPIRequest) SetImportProduct(_importProduct *ImportProduct) error {
	r._importProduct = _importProduct
	r.Set("import_product", _importProduct)
	return nil
}

// GetImportProduct ImportProduct Getter
func (r TaobaoFiveeImportproductPublishAPIRequest) GetImportProduct() *ImportProduct {
	return r._importProduct
}

var poolTaobaoFiveeImportproductPublishAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoFiveeImportproductPublishRequest()
	},
}

// GetTaobaoFiveeImportproductPublishRequest 从 sync.Pool 获取 TaobaoFiveeImportproductPublishAPIRequest
func GetTaobaoFiveeImportproductPublishAPIRequest() *TaobaoFiveeImportproductPublishAPIRequest {
	return poolTaobaoFiveeImportproductPublishAPIRequest.Get().(*TaobaoFiveeImportproductPublishAPIRequest)
}

// ReleaseTaobaoFiveeImportproductPublishAPIRequest 将 TaobaoFiveeImportproductPublishAPIRequest 放入 sync.Pool
func ReleaseTaobaoFiveeImportproductPublishAPIRequest(v *TaobaoFiveeImportproductPublishAPIRequest) {
	v.Reset()
	poolTaobaoFiveeImportproductPublishAPIRequest.Put(v)
}
