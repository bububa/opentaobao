package product

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// Alitriptravelelementssearch 商家元素搜索
// alitrip.travel.elements.search
//
// 提供商家维护的景点、酒店、餐饮等元素搜索
func Alitriptravelelementssearch(clt *core.SDKClient, req *product.AlitriptravelelementssearchAPIRequest, session string) (*product.AlitriptravelelementssearchAPIResponse, error) {
	var resp product.AlitriptravelelementssearchAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
