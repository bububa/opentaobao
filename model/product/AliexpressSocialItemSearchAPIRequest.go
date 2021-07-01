package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AliexpressSocialItemSearchAPIRequest
AE社交选品 API请求
aliexpress.social.item.search

AE社交选品,通过各种筛选条件对社交商品池进行筛选 */
type AliexpressSocialItemSearchAPIRequest struct {
	model.Params
	// 是否有视频
	_hasVideo bool
	// order by properties
	_orderBy string
	// 是否逆序
	_desc bool
	// page size
	_pageSize int64
	// 是否免邮
	_isShipFree bool
	// 佣金最大值
	_commissionRateMax string
	// 佣金最小值
	_commissionRateMin string
	// 物流时效
	_logisticsTime int64
	// 是否联盟商品
	_allianceItem bool
	// 类目ID
	_cateId int64
	// 页码
	_pageNo int64
	// 最低价格
	_minPrice string
	// 最高价格
	_maxPrice string
	// 搜索关键字
	_keyword string
	// shipTo国家
	_shipTo string
	// 评价分数
	_commentScore string
	// 币种
	_currency string
	// locale,格式为language+"_"+country
	_locale string
}

// New
