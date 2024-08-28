package simba

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSubwayAdgroupOfflineFind 查询单元离线多日汇总列表
// taobao.subway.adgroup.offline.find
//
// 查询单元离线列表
func TaobaoSubwayAdgroupOfflineFind(ctx context.Context, clt *core.SDKClient, req *simba.TaobaoSubwayAdgroupOfflineFindAPIRequest, resp *simba.TaobaoSubwayAdgroupOfflineFindAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
