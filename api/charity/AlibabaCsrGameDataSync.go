package charity

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/charity"
)

// AlibabaCsrGameDataSync 公益互动 外部游戏数据同步
// alibaba.csr.game.data.sync
//
// 公益互动 外部游戏数据同步
func AlibabaCsrGameDataSync(ctx context.Context, clt *core.SDKClient, req *charity.AlibabaCsrGameDataSyncAPIRequest, resp *charity.AlibabaCsrGameDataSyncAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
