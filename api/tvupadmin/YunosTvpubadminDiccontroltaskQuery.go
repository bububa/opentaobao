package tvupadmin

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// YunosTvpubadminDiccontroltaskQuery 停开服任务列表
// yunos.tvpubadmin.diccontroltask.query
//
// 牌照方对终端设备的停开服管理
func YunosTvpubadminDiccontroltaskQuery(ctx context.Context, clt *core.SDKClient, req *tvupadmin.YunosTvpubadminDiccontroltaskQueryAPIRequest, resp *tvupadmin.YunosTvpubadminDiccontroltaskQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
