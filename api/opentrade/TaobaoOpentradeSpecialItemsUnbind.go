package opentrade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/opentrade"
)

// Taobaoopentradespecialitemsunbind 专属下单场景商品解绑
// taobao.opentrade.special.items.unbind
//
// 专属下单场景商品解绑
func Taobaoopentradespecialitemsunbind(clt *core.SDKClient, req *opentrade.TaobaoopentradespecialitemsunbindAPIRequest, session string) (*opentrade.TaobaoopentradespecialitemsunbindAPIResponse, error) {
	var resp opentrade.TaobaoopentradespecialitemsunbindAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
