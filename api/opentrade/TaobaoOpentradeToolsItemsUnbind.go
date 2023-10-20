package opentrade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/opentrade"
)

// Taobaoopentradetoolsitemsunbind 交易开放商品解绑
// taobao.opentrade.tools.items.unbind
//
// 交易开放商品解绑
func Taobaoopentradetoolsitemsunbind(clt *core.SDKClient, req *opentrade.TaobaoopentradetoolsitemsunbindAPIRequest, session string) (*opentrade.TaobaoopentradetoolsitemsunbindAPIResponse, error) {
	var resp opentrade.TaobaoopentradetoolsitemsunbindAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
