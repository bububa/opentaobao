package jst

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jst"
)

// TaobaoJstSmsSignnameCreate 淘宝短信签名创建接口
// taobao.jst.sms.signname.create
//
// 聚石塔短信签名创建接口
func TaobaoJstSmsSignnameCreate(clt *core.SDKClient, req *jst.TaobaoJstSmsSignnameCreateAPIRequest, resp *jst.TaobaoJstSmsSignnameCreateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
