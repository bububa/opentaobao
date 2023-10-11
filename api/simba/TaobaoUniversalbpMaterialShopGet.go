package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoUniversalbpMaterialShopGet 获取店铺信息
// taobao.universalbp.material.shop.get
//
// 获取店铺信息
func TaobaoUniversalbpMaterialShopGet(clt *core.SDKClient, req *simba.TaobaoUniversalbpMaterialShopGetAPIRequest, session string) (*simba.TaobaoUniversalbpMaterialShopGetAPIResponse, error) {
	var resp simba.TaobaoUniversalbpMaterialShopGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
