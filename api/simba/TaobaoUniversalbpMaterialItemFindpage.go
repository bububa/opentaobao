package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoUniversalbpMaterialItemFindpage 分页查询商品信息
// taobao.universalbp.material.item.findpage
//
// 分页获取店铺内的商品列表
func TaobaoUniversalbpMaterialItemFindpage(clt *core.SDKClient, req *simba.TaobaoUniversalbpMaterialItemFindpageAPIRequest, session string) (*simba.TaobaoUniversalbpMaterialItemFindpageAPIResponse, error) {
	var resp simba.TaobaoUniversalbpMaterialItemFindpageAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
