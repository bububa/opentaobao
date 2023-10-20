package travel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/travel"
)

// TaobaoAlitripTravelItemSkuPriceModify 【API3.0】日期级别日历价格库存修改，增量维护
// taobao.alitrip.travel.item.sku.price.modify
//
// 【API3.0】日期级别日历价格库存增量维护
func TaobaoAlitripTravelItemSkuPriceModify(clt *core.SDKClient, req *travel.TaobaoAlitripTravelItemSkuPriceModifyAPIRequest, resp *travel.TaobaoAlitripTravelItemSkuPriceModifyAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
