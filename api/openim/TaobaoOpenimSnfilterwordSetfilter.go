package openim

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/openim"
)

// TaobaoOpenimSnfilterwordSetfilter 关键词过滤
// taobao.openim.snfilterword.setfilter
//
// 设置openim关键词过滤
func TaobaoOpenimSnfilterwordSetfilter(ctx context.Context, clt *core.SDKClient, req *openim.TaobaoOpenimSnfilterwordSetfilterAPIRequest, resp *openim.TaobaoOpenimSnfilterwordSetfilterAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
