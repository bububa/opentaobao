package icbu

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaIcbuProductUpdateAPIRequest
修改商品 API请求
alibaba.icbu.product.update

修改国际站商品，支持询盘商品和在线批发商品，支持英文商品和多语言商品 */
type AlibabaIcbuProductUpdateAPIRequest struct {
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
	// 商品类型，在线批发商品(wholesale)或者询盘商品(sourcing)
	_productType string
	// 询盘商品交易信息
	_sourcingTrade *SourcingTrade
	// 商品名称，最多128个字符
	_subject string
	// 在线批发商品交易信息
	_wholesaleTrade *WholesaleTrade
	// 发布的市场，支持main/onesite，默认main发到主市场，填onesite发布为商机通产品
	_market string
	// 智能编辑，不填写使用原来的。注意必须和详情的格式一致
	_isSmartEdit bool
	// 定制信息
	_customInfo *CustomInfo
	// 混淆商品ID
	_productId string
}

// New
