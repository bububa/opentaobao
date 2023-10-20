package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoUniversalbpMaterialShopGet 获取店铺信息
// taobao.universalbp.material.shop.get
//
// 获取店铺信息
func TaobaoUniversalbpMaterialShopGet(clt *core.SDKClient, req *simba.TaobaoUniversalbpMaterialShopGetAPIRequest, resp *simba.TaobaoUniversalbpMaterialShopGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
