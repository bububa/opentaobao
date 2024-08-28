package jst

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jst"
)

// TaobaoJstSmsSignnameModify 淘宝短信签名修改
// taobao.jst.sms.signname.modify
//
// 淘宝短信签名修改，只能修改还未被审核的签名。
func TaobaoJstSmsSignnameModify(ctx context.Context, clt *core.SDKClient, req *jst.TaobaoJstSmsSignnameModifyAPIRequest, resp *jst.TaobaoJstSmsSignnameModifyAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
