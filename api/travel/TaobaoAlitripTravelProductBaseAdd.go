package travel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/travel"
)

// TaobaoAlitripTravelProductBaseAdd 供应商新增产品API
// taobao.alitrip.travel.product.base.add
//
// 飞猪供销平台供应商可通过该API发布新产品
func TaobaoAlitripTravelProductBaseAdd(clt *core.SDKClient, req *travel.TaobaoAlitripTravelProductBaseAddAPIRequest, resp *travel.TaobaoAlitripTravelProductBaseAddAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
