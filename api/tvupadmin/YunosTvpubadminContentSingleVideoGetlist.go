package tvupadmin

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// YunosTvpubadminContentSingleVideoGetlist 单视频审核获取视频列表
// yunos.tvpubadmin.content.single.video.getlist
//
// 牌照方在审核单视频时获取单视频列表接口
func YunosTvpubadminContentSingleVideoGetlist(ctx context.Context, clt *core.SDKClient, req *tvupadmin.YunosTvpubadminContentSingleVideoGetlistAPIRequest, resp *tvupadmin.YunosTvpubadminContentSingleVideoGetlistAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
