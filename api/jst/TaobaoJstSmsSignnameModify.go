package jst

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jst"
)

// TaobaoJstSmsSignnameModify 淘宝短信签名修改
// taobao.jst.sms.signname.modify
//
// 淘宝短信签名修改，只能修改还未被审核的签名。
func TaobaoJstSmsSignnameModify(clt *core.SDKClient, req *jst.TaobaoJstSmsSignnameModifyAPIRequest, resp *jst.TaobaoJstSmsSignnameModifyAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
