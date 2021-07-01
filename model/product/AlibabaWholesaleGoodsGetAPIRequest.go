package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWholesaleGoodsGetAPIRequest
查询阿里巴巴批发市场商品详情 API请求
alibaba.wholesale.goods.get

查询阿里巴巴批发市场商品详情 */
type AlibabaWholesaleGoodsGetAPIRequest struct {
	model.Params
	// country_code
	_countryCode string
	// id
	_id string
}

// New
