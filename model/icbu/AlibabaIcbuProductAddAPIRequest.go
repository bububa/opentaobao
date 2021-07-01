package icbu

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaIcbuProductAddAPIRequest
发布产品 API请求
alibaba.icbu.product.add

发布商品,支持sourcing/一口价商品，支持英文和多种语言原发商品 */
type AlibabaIcbuProductAddAPIRequest struct {
	model.Params
	// 商品属性和属性值
	_attributes []ProductAttribute
	// 根据数量设置的折扣价
	_bulkDiscountPrices []BulkDiscountPrice
	// 类目ID
	_categoryId int64
	// 商品详情描述，可包含图片中心的图片URL
	_description string
	// 补充信息
	_extraContext string
	// 分组ID
	_groupId int64
	// 关键词，不要包含特殊符号（如,;），最多三个
	_keywords []string
	// 语种，参见FAQ 语种枚举值
	_language string
	// 商品主图
	_mainImage *MainImage
	// 商品SKU定义
	_productSku *ProductSku
	// 商品类型，在线批发商品(wholesale)或者询盘商品(sourcing)，值为wholesale时，必须填写wholesale_trade
	_productType string
	// 询盘商品交易信息
	_sourcingTrade *SourcingTrade
	// 商品名称，最多128个字符
	_subject string
	// 在线批发商品交易信息
	_wholesaleTrade *WholesaleTrade
	// 发布的市场，支持main，发到主市场
	_market string
	// 是否智能编辑，如果不传，默认为false
	_isSmartEdit bool
	// 定制信息
	_customInfo *CustomInfo
}

// New
