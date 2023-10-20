package jst

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jst"
)

// TaobaoJstSmsSignnameQuery 淘宝短信签名查询
// taobao.jst.sms.signname.query
//
// 淘宝短信签名查询
func TaobaoJstSmsSignnameQuery(clt *core.SDKClient, req *jst.TaobaoJstSmsSignnameQueryAPIRequest, resp *jst.TaobaoJstSmsSignnameQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
