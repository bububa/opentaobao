package alihealth2

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

// TaobaoDrugShopList 查询卖家外卖店列表
// taobao.drug.shop.list
//
// 查询卖家外卖店列表
func TaobaoDrugShopList(clt *core.SDKClient, req *alihealth2.TaobaoDrugShopListAPIRequest, resp *alihealth2.TaobaoDrugShopListAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
