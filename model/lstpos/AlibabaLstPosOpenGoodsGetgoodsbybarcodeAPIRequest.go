package lstpos

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLstPosOpenGoodsGetgoodsbybarcodeAPIRequest ISV条码库查询接口 API请求
// alibaba.lst.pos.open.goods.getgoodsbybarcode
//
// ISV条码库查询接口
type AlibabaLstPosOpenGoodsGetgoodsbybarcodeAPIRequest struct {
	model.Params
	// 商品条码
	_barcode string
}

// NewAlibabaLstPosOpenGoodsGetgoodsbybarcodeRequest 初始化AlibabaLstPosOpenGoodsGetgoodsbybarcodeAPIRequest对象
func NewAlibabaLstPosOpenGoodsGetgoodsbybarcodeRequest() *AlibabaLstPosOpenGoodsGetgoodsbybarcodeAPIRequest {
	return &AlibabaLstPosOpenGoodsGetgoodsbybarcodeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaLstPosOpenGoodsGetgoodsbybarcodeAPIRequest) GetApiMethodName() string {
	return "alibaba.lst.pos.open.goods.getgoodsbybarcode"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaLstPosOpenGoodsGetgoodsbybarcodeAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Barcode Setter
// 商品条码
func (r *AlibabaLstPosOpenGoodsGetgoodsbybarcodeAPIRequest) SetBarcode(_barcode string) error {
	r._barcode = _barcode
	r.Set("barcode", _barcode)
	return nil
}

// Get Barcode Getter
func (r AlibabaLstPosOpenGoodsGetgoodsbybarcodeAPIRequest) GetBarcode() string {
	return r._barcode
}
