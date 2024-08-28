package util

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/util"
)

// TaobaoAppipGet 获取ISV发起请求服务器IP
// taobao.appip.get
//
// 获取ISV发起请求服务器IP
func TaobaoAppipGet(ctx context.Context, clt *core.SDKClient, req *util.TaobaoAppipGetAPIRequest, resp *util.TaobaoAppipGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
