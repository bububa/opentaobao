package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkBmStockPublishAPIRequest
品牌营销涉及到的商品的库存同步接口 API请求
alibaba.wdk.bm.stock.publish

用于操作sku的库存 */
type AlibabaWdkBmStockPublishAPIRequest struct {
	model.Params
	// 批量入参
	_skuStockPublishParamList []SkuStockPublishParamDo
}

// New
