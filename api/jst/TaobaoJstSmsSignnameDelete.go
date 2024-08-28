package jst

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jst"
)

// TaobaoJstSmsSignnameDelete 淘宝短信签名删除
// taobao.jst.sms.signname.delete
//
// 淘宝短信签名删除
func TaobaoJstSmsSignnameDelete(ctx context.Context, clt *core.SDKClient, req *jst.TaobaoJstSmsSignnameDeleteAPIRequest, resp *jst.TaobaoJstSmsSignnameDeleteAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
