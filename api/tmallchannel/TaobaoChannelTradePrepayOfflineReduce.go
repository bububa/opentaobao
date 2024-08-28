package tmallchannel

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallchannel"
)

// TaobaoChannelTradePrepayOfflineReduce 渠道分销供应商上传线下流水预存款（减少）
// taobao.channel.trade.prepay.offline.reduce
//
// 渠道分销供应商上传线下流水预存款（减少）
func TaobaoChannelTradePrepayOfflineReduce(ctx context.Context, clt *core.SDKClient, req *tmallchannel.TaobaoChannelTradePrepayOfflineReduceAPIRequest, resp *tmallchannel.TaobaoChannelTradePrepayOfflineReduceAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
