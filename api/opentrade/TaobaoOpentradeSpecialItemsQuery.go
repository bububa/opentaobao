package opentrade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/opentrade"
)

// Taobaoopentradespecialitemsquery 专属下单获取商品绑定信息
// taobao.opentrade.special.items.query
//
// 专属下单获取商品绑定信息
func Taobaoopentradespecialitemsquery(clt *core.SDKClient, req *opentrade.TaobaoopentradespecialitemsqueryAPIRequest, session string) (*opentrade.TaobaoopentradespecialitemsqueryAPIResponse, error) {
	var resp opentrade.TaobaoopentradespecialitemsqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
