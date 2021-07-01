package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkSkuBarcodeQueryAPIRequest
商品条码查询接口 API请求
alibaba.wdk.sku.barcode.query

查询商品编码，支持一品多码 */
type AlibabaWdkSkuBarcodeQueryAPIRequest struct {
	model.Params
	// 商品编码
	_skuCode string
}

// New
