package gameact

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/gameact"
)

// TaobaoDeActivityLuckydraw 抽奖
// taobao.de.activity.luckydraw
//
// 用于激励平台对外提供抽奖功能，包括但不限于集分宝、红包、宝点、淘金币、淘彩票等
func TaobaoDeActivityLuckydraw(ctx context.Context, clt *core.SDKClient, req *gameact.TaobaoDeActivityLuckydrawAPIRequest, resp *gameact.TaobaoDeActivityLuckydrawAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
