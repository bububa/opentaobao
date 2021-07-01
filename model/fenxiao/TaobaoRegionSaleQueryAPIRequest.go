package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoRegionSaleQueryAPIRequest
查询商品销售区域 API请求
taobao.region.sale.query

查询商品销售区域 */
type TaobaoRegionSaleQueryAPIRequest struct {
	model.Params
	// 商品id
	_itemId int64
	// 无sku传0
	_skuId int64
	// 1:国家;2:省;3: 市;4:区县
	_saleRegionLevel int64
}

// New
