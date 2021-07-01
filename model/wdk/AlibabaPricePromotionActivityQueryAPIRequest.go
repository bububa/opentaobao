package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaPricePromotionActivityQueryAPIRequest
查询盒马帮档期活动详情 API请求
alibaba.price.promotion.activity.query

查询盒马帮档期活动详情 */
type AlibabaPricePromotionActivityQueryAPIRequest struct {
	model.Params
	// 页码
	_page int64
	// 外部档期code
	_outerPromotionCode string
	// TOB店仓编码
	_ouCode string
	// 页码大小
	_pageSize int64
}

// New
