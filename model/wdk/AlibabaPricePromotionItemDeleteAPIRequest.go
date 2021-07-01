package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaPricePromotionItemDeleteAPIRequest
批量删除档期 API请求
alibaba.price.promotion.item.delete

盒马帮批量删除档期商品 */
type AlibabaPricePromotionItemDeleteAPIRequest struct {
	model.Params
	// 商品code
	_skuCodes []string
	// toB渠道店OU
	_ouCode string
	// 外部档期编码
	_outerPromotionCode string
	// 盒马唯一标识
	_uniqueId string
}

// New
