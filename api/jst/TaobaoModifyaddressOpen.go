package jst

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jst"
)

// TaobaoModifyaddressOpen 淘宝自助修改地址服务开通
// taobao.modifyaddress.open
//
// 商家自助修改地址服务开通
func TaobaoModifyaddressOpen(clt *core.SDKClient, req *jst.TaobaoModifyaddressOpenAPIRequest, resp *jst.TaobaoModifyaddressOpenAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
