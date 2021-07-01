package lstpos

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaLstPosOpenGoodsGetgoodsbybarcodeAPIRequest
ISV条码库查询接口 API请求
alibaba.lst.pos.open.goods.getgoodsbybarcode

ISV条码库查询接口 */
type AlibabaLstPosOpenGoodsGetgoodsbybarcodeAPIRequest struct {
	model.Params
	// 商品条码
	_barcode string
}

// New
