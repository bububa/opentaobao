package lstpos

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabalstposopengoodsgetgoodsbybarcodeAPIRequest ISV条码库查询接口 API请求
// alibaba.lst.pos.open.goods.getgoodsbybarcode
//
// ISV条码库查询接口
type AlibabalstposopengoodsgetgoodsbybarcodeAPIRequest struct {
	model.Params
	// 商品条码
	_barcode string
}

// NewAlibabalstposopengoodsgetgoodsbybarcodeRequest 初始化AlibabalstposopengoodsgetgoodsbybarcodeAPIRequest对象
func NewAlibabalstposopengoodsgetgoodsbybarcodeRequest() *AlibabalstposopengoodsgetgoodsbybarcodeAPIRequest {
	return &AlibabalstposopengoodsgetgoodsbybarcodeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabalstposopengoodsgetgoodsbybarcodeAPIRequest) GetApiMethodName() string {
	return "alibaba.lst.pos.open.goods.getgoodsbybarcode"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabalstposopengoodsgetgoodsbybarcodeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabalstposopengoodsgetgoodsbybarcodeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBarcode is Barcode Setter
// 商品条码
func (r *AlibabalstposopengoodsgetgoodsbybarcodeAPIRequest) SetBarcode(_barcode string) error {
	r._barcode = _barcode
	r.Set("barcode", _barcode)
	return nil
}

// GetBarcode Barcode Getter
func (r AlibabalstposopengoodsgetgoodsbybarcodeAPIRequest) GetBarcode() string {
	return r._barcode
}
