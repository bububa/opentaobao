package lstpos

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaLstPosOpenGoodsGetgoodsbybarcodeAPIRequest) Reset() {
	r._barcode = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaLstPosOpenGoodsGetgoodsbybarcodeAPIRequest) GetApiMethodName() string {
	return "alibaba.lst.pos.open.goods.getgoodsbybarcode"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaLstPosOpenGoodsGetgoodsbybarcodeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaLstPosOpenGoodsGetgoodsbybarcodeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBarcode is Barcode Setter
// 商品条码
func (r *AlibabaLstPosOpenGoodsGetgoodsbybarcodeAPIRequest) SetBarcode(_barcode string) error {
	r._barcode = _barcode
	r.Set("barcode", _barcode)
	return nil
}

// GetBarcode Barcode Getter
func (r AlibabaLstPosOpenGoodsGetgoodsbybarcodeAPIRequest) GetBarcode() string {
	return r._barcode
}

var poolAlibabaLstPosOpenGoodsGetgoodsbybarcodeAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaLstPosOpenGoodsGetgoodsbybarcodeRequest()
	},
}

// GetAlibabaLstPosOpenGoodsGetgoodsbybarcodeRequest 从 sync.Pool 获取 AlibabaLstPosOpenGoodsGetgoodsbybarcodeAPIRequest
func GetAlibabaLstPosOpenGoodsGetgoodsbybarcodeAPIRequest() *AlibabaLstPosOpenGoodsGetgoodsbybarcodeAPIRequest {
	return poolAlibabaLstPosOpenGoodsGetgoodsbybarcodeAPIRequest.Get().(*AlibabaLstPosOpenGoodsGetgoodsbybarcodeAPIRequest)
}

// ReleaseAlibabaLstPosOpenGoodsGetgoodsbybarcodeAPIRequest 将 AlibabaLstPosOpenGoodsGetgoodsbybarcodeAPIRequest 放入 sync.Pool
func ReleaseAlibabaLstPosOpenGoodsGetgoodsbybarcodeAPIRequest(v *AlibabaLstPosOpenGoodsGetgoodsbybarcodeAPIRequest) {
	v.Reset()
	poolAlibabaLstPosOpenGoodsGetgoodsbybarcodeAPIRequest.Put(v)
}
