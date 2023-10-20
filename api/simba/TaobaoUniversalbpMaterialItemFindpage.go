package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoUniversalbpMaterialItemFindpage 分页查询商品信息
// taobao.universalbp.material.item.findpage
//
// 分页获取店铺内的商品列表
func TaobaoUniversalbpMaterialItemFindpage(clt *core.SDKClient, req *simba.TaobaoUniversalbpMaterialItemFindpageAPIRequest, resp *simba.TaobaoUniversalbpMaterialItemFindpageAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
