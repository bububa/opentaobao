package shop

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/shop"
)

// TaobaoShopUpdate 更新店铺基本信息
// taobao.shop.update
//
// 目前只支持标题、公告和描述的更新
func TaobaoShopUpdate(clt *core.SDKClient, req *shop.TaobaoShopUpdateAPIRequest, resp *shop.TaobaoShopUpdateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
