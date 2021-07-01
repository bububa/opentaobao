package travel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripTravelGereralproductUpdateAPIRequest
通用类目产品发布编辑 API请求
alitrip.travel.gereralproduct.update

提供给飞猪供销平台供应商发布编辑通用类目产品的API */
type AlitripTravelGereralproductUpdateAPIRequest struct {
	model.Params
	// 产品基本信息
	_baseInfo *GeneralProductBaseInfo
	// 退款规则结构
	_refundInfo *ItemRefundInfo
	// 必填，预定规则结构。示例： [{ "rule_type": "fee_excluded", "rule_desc": "费用包含描述"},{ "rule_type": "fee_included", "rule_desc": "费用不含描述"},{ "rule_type": "order_info", "rule_desc": "预定须知描述"}]
	_bookingRules []BookingRuleInfo
	// 产品销售信息
	_productSaleInfo *ProductSaleInfo
	// 更新sku信息，仅限日历商品使用
	_dateSkuInfoList []DateSkuInfo
}

// New
