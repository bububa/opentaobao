package alihealth2

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

// TaobaoDrugShopList 查询卖家外卖店列表
// taobao.drug.shop.list
//
// 查询卖家外卖店列表
func TaobaoDrugShopList(ctx context.Context, clt *core.SDKClient, req *alihealth2.TaobaoDrugShopListAPIRequest, resp *alihealth2.TaobaoDrugShopListAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
