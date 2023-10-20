package opentrade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/opentrade"
)

// Taobaoopentradetoolsitemsquery 交易开放获取商品绑定信息
// taobao.opentrade.tools.items.query
//
// 交易开放获取商品绑定信息
func Taobaoopentradetoolsitemsquery(clt *core.SDKClient, req *opentrade.TaobaoopentradetoolsitemsqueryAPIRequest, session string) (*opentrade.TaobaoopentradetoolsitemsqueryAPIResponse, error) {
	var resp opentrade.TaobaoopentradetoolsitemsqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
