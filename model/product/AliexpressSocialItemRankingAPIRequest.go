package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AliexpressSocialItemRankingAPIRequest
社交排行榜 API请求
aliexpress.social.item.ranking

社交商品成交排行榜 */
type AliexpressSocialItemRankingAPIRequest struct {
	model.Params
	// 币种
	_currency string
	// 国家列表
	_countryList []string
	// locale，格式为language+"_"+country
	_locale string
	// 页码
	_pageNo int64
	// 类目ID
	_cateId int64
	// 每页条数
	_pageSize int64
}

// New
