package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkSkuBarcodeQuery 商品条码查询接口
// alibaba.wdk.sku.barcode.query
//
// 查询商品编码，支持一品多码
func AlibabaWdkSkuBarcodeQuery(clt *core.SDKClient, req *wdk.AlibabaWdkSkuBarcodeQueryAPIRequest, resp *wdk.AlibabaWdkSkuBarcodeQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
