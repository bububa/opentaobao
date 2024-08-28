package tvupadmin

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// YunosTvpubadminDiccontroltaskGetinfo 获取停开服任务详情
// yunos.tvpubadmin.diccontroltask.getinfo
//
// 获取停开服任务详情
func YunosTvpubadminDiccontroltaskGetinfo(ctx context.Context, clt *core.SDKClient, req *tvupadmin.YunosTvpubadminDiccontroltaskGetinfoAPIRequest, resp *tvupadmin.YunosTvpubadminDiccontroltaskGetinfoAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
