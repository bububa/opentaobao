package tvupadmin

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// YunosTvpubadminContentVideoGetauditlist 迎客松视频审核记录查询
// yunos.tvpubadmin.content.video.getauditlist
//
// 迎客松视频审核记录查询
func YunosTvpubadminContentVideoGetauditlist(ctx context.Context, clt *core.SDKClient, req *tvupadmin.YunosTvpubadminContentVideoGetauditlistAPIRequest, resp *tvupadmin.YunosTvpubadminContentVideoGetauditlistAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
