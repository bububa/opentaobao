package opentrade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/opentrade"
)

// Taobaoopentradespecialitemsbind 专属下单场景商品绑定
// taobao.opentrade.special.items.bind
//
// 专属下单场景商品绑定
func Taobaoopentradespecialitemsbind(clt *core.SDKClient, req *opentrade.TaobaoopentradespecialitemsbindAPIRequest, session string) (*opentrade.TaobaoopentradespecialitemsbindAPIResponse, error) {
	var resp opentrade.TaobaoopentradespecialitemsbindAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
