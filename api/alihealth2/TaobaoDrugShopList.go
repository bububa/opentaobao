package alihealth2

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

// Taobaodrugshoplist 查询卖家外卖店列表
// taobao.drug.shop.list
//
// 查询卖家外卖店列表
func Taobaodrugshoplist(clt *core.SDKClient, req *alihealth2.TaobaodrugshoplistAPIRequest, session string) (*alihealth2.TaobaodrugshoplistAPIResponse, error) {
	var resp alihealth2.TaobaodrugshoplistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
