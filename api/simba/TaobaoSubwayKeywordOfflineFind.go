package simba

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSubwayKeywordOfflineFind 查询关键词离线多日汇总报表
// taobao.subway.keyword.offline.find
//
// 获取关键词离线报表
func TaobaoSubwayKeywordOfflineFind(ctx context.Context, clt *core.SDKClient, req *simba.TaobaoSubwayKeywordOfflineFindAPIRequest, resp *simba.TaobaoSubwayKeywordOfflineFindAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
