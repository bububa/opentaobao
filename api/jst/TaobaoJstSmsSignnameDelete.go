package jst

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jst"
)

// TaobaoJstSmsSignnameDelete 淘宝短信签名删除
// taobao.jst.sms.signname.delete
//
// 淘宝短信签名删除
func TaobaoJstSmsSignnameDelete(clt *core.SDKClient, req *jst.TaobaoJstSmsSignnameDeleteAPIRequest, resp *jst.TaobaoJstSmsSignnameDeleteAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
