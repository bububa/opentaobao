package jst

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jst"
)

// TaobaoModifyaddressOpen 淘宝自助修改地址服务开通
// taobao.modifyaddress.open
//
// 商家自助修改地址服务开通
func TaobaoModifyaddressOpen(ctx context.Context, clt *core.SDKClient, req *jst.TaobaoModifyaddressOpenAPIRequest, resp *jst.TaobaoModifyaddressOpenAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
