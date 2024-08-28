package tmallchannel

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallchannel"
)

// TaobaoChannelTradePrepayOfflineAdd 渠道分销供应商上传线下流水预存款（增加）
// taobao.channel.trade.prepay.offline.add
//
// 渠道分销供应商上传线下流水预存款（增加）
func TaobaoChannelTradePrepayOfflineAdd(ctx context.Context, clt *core.SDKClient, req *tmallchannel.TaobaoChannelTradePrepayOfflineAddAPIRequest, resp *tmallchannel.TaobaoChannelTradePrepayOfflineAddAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
