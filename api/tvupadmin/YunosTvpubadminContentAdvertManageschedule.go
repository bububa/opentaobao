package tvupadmin

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// YunosTvpubadminContentAdvertManageschedule 广告牌照管控修改
// yunos.tvpubadmin.content.advert.manageschedule
//
// 广告牌照管控修改
func YunosTvpubadminContentAdvertManageschedule(ctx context.Context, clt *core.SDKClient, req *tvupadmin.YunosTvpubadminContentAdvertManagescheduleAPIRequest, resp *tvupadmin.YunosTvpubadminContentAdvertManagescheduleAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
