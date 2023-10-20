package opentrade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/opentrade"
)

// Taobaoopentradetoolsitemsbind 交易开放商品绑定
// taobao.opentrade.tools.items.bind
//
// 交易开放商品绑定
func Taobaoopentradetoolsitemsbind(clt *core.SDKClient, req *opentrade.TaobaoopentradetoolsitemsbindAPIRequest, session string) (*opentrade.TaobaoopentradetoolsitemsbindAPIResponse, error) {
	var resp opentrade.TaobaoopentradetoolsitemsbindAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
