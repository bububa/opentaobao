package travel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/travel"
)

// Taobaoalitriptravelitemskupackagemodify 【API3.0】套餐级别日历价格库存增删操作
// taobao.alitrip.travel.item.sku.package.modify
//
// 【API3.0】套餐级别日历价格库存增删操作
func Taobaoalitriptravelitemskupackagemodify(clt *core.SDKClient, req *travel.TaobaoalitriptravelitemskupackagemodifyAPIRequest, session string) (*travel.TaobaoalitriptravelitemskupackagemodifyAPIResponse, error) {
	var resp travel.TaobaoalitriptravelitemskupackagemodifyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
