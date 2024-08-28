package tvupadmin

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// YunosTvpubadminContentShowGetbyshowid 迎客松根据节目id获取节目元数据
// yunos.tvpubadmin.content.show.getbyshowid
//
// 迎客松根据节目id获取节目元数据
func YunosTvpubadminContentShowGetbyshowid(ctx context.Context, clt *core.SDKClient, req *tvupadmin.YunosTvpubadminContentShowGetbyshowidAPIRequest, resp *tvupadmin.YunosTvpubadminContentShowGetbyshowidAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
