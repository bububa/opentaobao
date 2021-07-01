package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWholesaleGoodsSearchAPIRequest
批发市场产品搜索 API请求
alibaba.wholesale.goods.search

批发市场产品搜索 */
type AlibabaWholesaleGoodsSearchAPIRequest struct {
	model.Params
	// SearchGoodsOption
	_paramSearchGoodsOption *SearchGoodsOption
}

// New
