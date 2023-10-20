package travel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/travel"
)

// TaobaoAlitripTravelProductSkuOverride （供销）产品级别日历价格库存修改，全量覆盖
// taobao.alitrip.travel.product.sku.override
//
// （供销）产品级别日历价格库存修改，全量覆盖
func TaobaoAlitripTravelProductSkuOverride(clt *core.SDKClient, req *travel.TaobaoAlitripTravelProductSkuOverrideAPIRequest, resp *travel.TaobaoAlitripTravelProductSkuOverrideAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
