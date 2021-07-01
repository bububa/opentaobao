package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkBmPaiyangStockQueryAPIRequest
派样商品门店库存查询接口 API请求
alibaba.wdk.bm.paiyang.stock.query

淘鲜达接入第三方进行派样，第三方查询派样商品的门店库存信息。 */
type AlibabaWdkBmPaiyangStockQueryAPIRequest struct {
	model.Params
	// 请求入参
	_isvShopStockParam *IsvShopStockParam
}

// New
