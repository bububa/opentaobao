package product

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// AlitripTravelElementsSearch 商家元素搜索
// alitrip.travel.elements.search
//
// 提供商家维护的景点、酒店、餐饮等元素搜索
func AlitripTravelElementsSearch(clt *core.SDKClient, req *product.AlitripTravelElementsSearchAPIRequest, resp *product.AlitripTravelElementsSearchAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
