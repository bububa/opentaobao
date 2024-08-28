package tvupadmin

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// YunosTvpubadminDataQuery 魔盒统计数据查询接口
// yunos.tvpubadmin.data.query
//
// 用于华数查询魔盒上的一些用户统计数据
func YunosTvpubadminDataQuery(ctx context.Context, clt *core.SDKClient, req *tvupadmin.YunosTvpubadminDataQueryAPIRequest, resp *tvupadmin.YunosTvpubadminDataQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
