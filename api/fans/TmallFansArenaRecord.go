package fans

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fans"
)

// TmallFansArenaRecord 记录完成擂台的用户
// tmall.fans.arena.record
//
// 记录完成擂台的用户和完成分数
func TmallFansArenaRecord(ctx context.Context, clt *core.SDKClient, req *fans.TmallFansArenaRecordAPIRequest, resp *fans.TmallFansArenaRecordAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
