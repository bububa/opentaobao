package jst

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jst"
)

// TaobaoOcApContracturlGet 按用户获取支付宝代扣协议链接地址
// taobao.oc.ap.contracturl.get
//
// 按用户获取支付宝代扣协议链接地址
func TaobaoOcApContracturlGet(ctx context.Context, clt *core.SDKClient, req *jst.TaobaoOcApContracturlGetAPIRequest, resp *jst.TaobaoOcApContracturlGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
