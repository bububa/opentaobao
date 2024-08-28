package tbtrade

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbtrade"
)

// TaobaoTopOaidMerge OAID订单合并
// taobao.top.oaid.merge
//
// 基于OAID（收件人ID， Open Addressee ID)做订单合并，确保相同收件人信息的订单合并到相同组。
func TaobaoTopOaidMerge(ctx context.Context, clt *core.SDKClient, req *tbtrade.TaobaoTopOaidMergeAPIRequest, resp *tbtrade.TaobaoTopOaidMergeAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
