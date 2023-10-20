package interact

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/interact"
)

// AlibabaInteractShopFavor 店铺收藏
// alibaba.interact.shop.favor
//
// 店铺收藏mtop接口开放鉴权接口，无入参出参，无安全风险，mtop接口开发 酒仙。
func AlibabaInteractShopFavor(clt *core.SDKClient, req *interact.AlibabaInteractShopFavorAPIRequest, resp *interact.AlibabaInteractShopFavorAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
