package tvupadmin

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// YunosTvpubadminDiccontroltaskUpdate 停开服任务状态变更
// yunos.tvpubadmin.diccontroltask.update
//
// 停开服任务状态变更
func YunosTvpubadminDiccontroltaskUpdate(ctx context.Context, clt *core.SDKClient, req *tvupadmin.YunosTvpubadminDiccontroltaskUpdateAPIRequest, resp *tvupadmin.YunosTvpubadminDiccontroltaskUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
