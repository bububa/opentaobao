package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripTravelElementsSearchAPIRequest
商家元素搜索 API请求
alitrip.travel.elements.search

提供商家维护的景点、酒店、餐饮等元素搜索 */
type AlitripTravelElementsSearchAPIRequest struct {
	model.Params
	// 商家id
	_sellerId int64
	// 查询关键词
	_query string
	// 查询数量，限制100
	_count int64
	// 资源类型
	_type int64
}

// New
