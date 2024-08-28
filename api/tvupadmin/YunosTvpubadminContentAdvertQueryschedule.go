package tvupadmin

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// YunosTvpubadminContentAdvertQueryschedule 广告牌照管控查询
// yunos.tvpubadmin.content.advert.queryschedule
//
// 广告牌照管控查询
func YunosTvpubadminContentAdvertQueryschedule(ctx context.Context, clt *core.SDKClient, req *tvupadmin.YunosTvpubadminContentAdvertQueryscheduleAPIRequest, resp *tvupadmin.YunosTvpubadminContentAdvertQueryscheduleAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
