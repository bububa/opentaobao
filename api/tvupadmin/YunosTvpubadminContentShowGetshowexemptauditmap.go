package tvupadmin

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// YunosTvpubadminContentShowGetshowexemptauditmap 迎客松批量查询节目某个牌照的免审状态
// yunos.tvpubadmin.content.show.getshowexemptauditmap
//
// 迎客松批量查询节目某个牌照的免审状态
func YunosTvpubadminContentShowGetshowexemptauditmap(ctx context.Context, clt *core.SDKClient, req *tvupadmin.YunosTvpubadminContentShowGetshowexemptauditmapAPIRequest, resp *tvupadmin.YunosTvpubadminContentShowGetshowexemptauditmapAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
