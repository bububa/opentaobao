package travel

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/travel"
)

// TaobaoAlitripTravelItemSkuOverride 【API3.0】商品级别日历价格库存修改，全量覆盖
// taobao.alitrip.travel.item.sku.override
//
// 旅行度假新商品日历价格库存信息修改接口 第三版。提供商家通过TOP API方式修改商品sku信息。
func TaobaoAlitripTravelItemSkuOverride(ctx context.Context, clt *core.SDKClient, req *travel.TaobaoAlitripTravelItemSkuOverrideAPIRequest, resp *travel.TaobaoAlitripTravelItemSkuOverrideAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
